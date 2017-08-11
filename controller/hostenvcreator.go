package controller

import (
	"os"
	"os/exec"
)

/*
To Create Clone Files of Source Code to Move the Code to the Container
and Keep and Copy in the Host Env
*/
func CreateRespectiveEnvOfLanguage(language_id int, path string, defualt_template_env_path string) (string, string) {
	//Create a Folder Same as the Repository ID
	folderName := CreateContainerExecEnv()
	//Construction of Dynamically Generated Repo Path
	path = path + "/" + folderName + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
	//Language ID Received must be according to the stored index in the Constants
	language_Indexes_Slice := ReturnContantLanguageid()
	//Contructing the Template Files of Particular Language
	language_template := defualt_template_env_path + "/" + language_Indexes_Slice[language_id] + "/*"
	//Copying the Language Template to the HostEnv for Wrting Code by the user
	if _, err := exec.Command("/bin/bash", "-c", "cp "+language_template+" "+path).Output(); err != nil {
		panic(err)
	}

	// The folderName is also Resp.iD
	return path, folderName
}
