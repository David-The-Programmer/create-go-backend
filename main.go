package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	modfile "golang.org/x/mod/modfile"
)

func main() {
	// TODO: Validation of input from user
	scanner := bufio.NewScanner(os.Stdin)
	// TODO: Need to account for relative path inputs, user current working dir
	fmt.Printf("Enter your project folder path: ")
	projectPath := ""
	if scanner.Scan() {
		projectPath = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("Failed to read project folder path: %w", err)
		slog.Error(err.Error())
		return
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("Failed to find user home dir: %w", err)
		slog.Error(err.Error())
		return
	}
	projectPath = strings.Replace(projectPath, "~", userHomeDir, 1)

	fmt.Printf("Enter your go module path: ")
	goModulePath := ""
	if scanner.Scan() {
		goModulePath = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("Failed to read go module path: %w", err)
		slog.Error(err.Error())
		return
	}

	fmt.Printf("Enter the go version of your go module: ")
	goVersion := ""
	if scanner.Scan() {
		goVersion = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		err = fmt.Errorf("Failed to read go version: %w", err)
		slog.Error(err.Error())
		return
	}
	// TODO: Pipe slog to log errors to file instead

	fmt.Printf("Creating project folder %s...\n", projectPath)
	err = os.Mkdir(projectPath, 0750)
	if err != nil {
		err = fmt.Errorf("Failed to create project folder: %w", err)
		slog.Error(err.Error())
		return
	}

	fmt.Println("Fetching download links of files in repo template folder...")
	ghAPIDomain := "https://api.github.com"
	owner := "David-The-Programmer"
	repo := "create-go-backend"
	path := "template"
	queryParam := "ref=main"
	repoTemplateContentURL := fmt.Sprintf("%s/repos/%s/%s/contents/%s?%s", ghAPIDomain, owner, repo, path, queryParam)
	urls, err := fileDownloadURLs(repoTemplateContentURL)
	if err != nil {
		err = fmt.Errorf("Failed to get all download URLs of files of repo template folder: %w", err)
		slog.Error(err.Error())
		return
	}

	fmt.Printf("Downloading files in template folder into %s\n", projectPath)
	for _, url := range urls {
		i := strings.Index(url, path)
		fileName := url[i+len(path) : len(url)]
		projectFilepath := filepath.Join(projectPath, fileName)
		fmt.Printf("Downloading file %s into %s\n", fileName, projectFilepath)
		err := downloadFile(url, projectFilepath)
		if err != nil {
			err = fmt.Errorf("Failed to download file %s: %w", fileName, err)
			slog.Error(err.Error())
			return
		}
	}
	fmt.Println("Download complete!")

	fmt.Printf("Initalising Go module: %s\n", goModulePath)
	goModFilepath := filepath.Join(projectPath, "go.mod")
	_, err = os.Create(goModFilepath)
	if err != nil {
		err = fmt.Errorf("Failed to create go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	goModFile, err := modfile.Parse(goModFilepath, []byte{}, nil)
	if err != nil {
		err = fmt.Errorf("Failed to parse go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	err = goModFile.AddModuleStmt(goModulePath)
	if err != nil {
		err = fmt.Errorf("Failed to add go module path to go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	err = goModFile.AddGoStmt(goVersion)
	if err != nil {
		err = fmt.Errorf("Failed to add go version to go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	goModFileContent, err := goModFile.Format()
	if err != nil {
		err = fmt.Errorf("Failed to format go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	err = os.WriteFile(goModFilepath, goModFileContent, 0660)
	if err != nil {
		err = fmt.Errorf("Failed to write to go.mod file: %w", err)
		slog.Error(err.Error())
		return
	}
	fmt.Println("Go module initialised!")

	fmt.Println("Creating .env file...")
	envFilepath := filepath.Join(projectPath, ".env")
	err = os.WriteFile(envFilepath, []byte(fmt.Sprintf("GO_VERSION=%s", goVersion)), 0660)
	if err != nil {
		err = fmt.Errorf("Failed to write to .env file: %w", err)
		slog.Error(err.Error())
		return
	}
	fmt.Println("Creation of Go backend project folder complete!")
}

// See https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-repository-content
func fileDownloadURLs(repoContentURL string) ([]string, error) {
	client := &http.Client{}
	// TODO: Add cancellation of request after timeout
	req, err := http.NewRequest(http.MethodGet, repoContentURL, nil)
	if err != nil {
		err = fmt.Errorf("Failed to make request to retrieve repo content info: %w", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("Failed to send request to retrieve repo content info: %w", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Failed to read repo content info: %w", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("Failed to retrieve repo content info, status: %s, body: %s ", resp.Status, body)
		return nil, err
	}
	type repoContentInfoLinks struct {
		Self string `json:self`
		Git  string `json:git`
		HTML string `json:html`
	}
	type repoContentInfo struct {
		Name        string               `json:"name"`
		Path        string               `json:"path"`
		SHA         string               `json:"sha"`
		Size        int                  `json:"size"`
		URL         string               `json:"url"`
		HTMLURL     string               `json:"html_url"`
		GitURL      string               `json:"git_url"`
		DownloadURL string               `json:"download_url"`
		Type        string               `json:"type"`
		Links       repoContentInfoLinks `json:"links"`
	}
	contentInfo := []repoContentInfo{}
	err = json.Unmarshal(body, &contentInfo)
	if err != nil {
		err = fmt.Errorf("Failed to unmarshal repo content info: %w", err)
		return nil, err
	}
	downloadURLs := []string{}
	for _, content := range contentInfo {
		if content.Type == "dir" {
			urls, err := fileDownloadURLs(content.URL)
			if err != nil {
				return nil, err
			}
			downloadURLs = append(downloadURLs, urls...)
			continue
		}
		downloadURLs = append(downloadURLs, content.DownloadURL)
	}
	return downloadURLs, nil
}

func downloadFile(url string, saveFilepath string) error {
	downloadReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("Failed to make request to download file contents: %w", err)
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(downloadReq)
	if err != nil {
		err = fmt.Errorf("Failed to send request to download file contents: %w", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to download file contents, status: %s", resp.Status)
		return err
	}
	nestedDirsPath := filepath.Dir(saveFilepath)
	err = os.MkdirAll(nestedDirsPath, 0750)
	if err != nil {
		err = fmt.Errorf("Failed to create nested dirs %s: %w", nestedDirsPath, err)
		return err
	}
	file, err := os.Create(saveFilepath)
	if err != nil {
		err = fmt.Errorf("Failed to create local file to copy downloaded file contents to: %w", err)
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		err = fmt.Errorf("Failed to copy downloaded file contents to local file: %w", err)
		return err
	}
	return nil
}
