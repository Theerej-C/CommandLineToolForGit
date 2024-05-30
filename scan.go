package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strings"
)

func scan(folder string){
	fmt.Printf("Found folders:\n\n")
	repo := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsTofile(filePath,repo)
	fmt.Printf(("\n\nSuccessfully added\n\n"))
}

func addNewSliceElementsTofile(filePath string, repo []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(repo,existingRepos)
	dumpStringSliceToFIle(repos,filePath)
}

func dumpStringSliceToFIle(repos []string, filePath string) {
	content := strings.Join(repos,"\n")
	os.WriteFile(filePath,[]byte(content),0755)
}

func joinSlices(new []string, existingRepos []string) []string{
	for _,val := range new{
		if !sliceContains(existingRepos,val){
			existingRepos = append(existingRepos, val)
		}
	}
	return existingRepos
}

func sliceContains(existingRepos []string, value string) bool{
	for _,val := range existingRepos{
		if val == value{
			return true
		}
	}
	return false
}

func parseFileLinesToSlice(filePath string) []string{
	f:= openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		lines= append(lines, scanner.Text())
	}
	if err := scanner.Err() ; err != nil{
		if err != io.EOF{
			panic(err)
		}
	}
	return lines
}

func openFile(filePath string) *os.File{
	f,err := os.OpenFile(filePath,os.O_APPEND|os.O_RDWR,0755)
	if err != nil{
		if os.IsNotExist(err){
			_,err = os.Create(filePath)
			if err != nil{
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	return f
}

func getDotFilePath() string{
	usr,err := user.Current()
	if err != nil{
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "/.gogitlocalstats"
	return dotFile
}

func recursiveScanFolder(folder string) []string{
	return scanGitFolder(make([]string,0),folder)
}
func scanGitFolder(folders []string,folder string) []string{
	folder  = strings.TrimSuffix(folder,"/")
	f,err := os.Open(folder)
	if err != nil{
		log.Fatal(err)
	}
	files, err := f.Readdir(-1);
	f.Close()
	if err != nil{
		log.Fatal(err)
	}
	var path string
	for _,file := range files{
		if file.IsDir(){
			path = folder + "/" + file.Name()
			if file.Name() == ".git"{
				path = strings.TrimSuffix(path,"/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules"{
				continue
			}
			folders = scanGitFolder(folders,path)
		}
	}
	return folders
}
