package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/furikuri/webhook"
	"os"
)

func main2() {
	main2()
	webhook.Server(func(w webhook.Webhook) webhook.Verification {
		return webhook.Verification{
			State:       "S",
			Description: "D",
			Context:     "C",
		}
	})
}

func main() {
	client, _ := docker.NewClientFromEnv()
	client.StopContainer("hello-docker", 30)
	client.RemoveContainer(docker.RemoveContainerOptions{
		ID:            "hello-docker",
		RemoveVolumes: false,
		Force:         false,
	})

	exposedCadvPort := map[docker.Port]struct{}{
		"8000/tcp": {}}

	createContConf := docker.Config{
		ExposedPorts: exposedCadvPort,
		Image:        "furikuri/hello-docker"}

	portBindings := map[docker.Port][]docker.PortBinding{
		"8000/tcp": []docker.PortBinding{docker.PortBinding{HostPort: "8000"}}}

	createContHostConfig := docker.HostConfig{
		PortBindings:    portBindings,
		PublishAllPorts: true,
		Privileged:      false}

	createContOps := docker.CreateContainerOptions{
		Name:       "hello-docker",
		Config:     &createContConf,
		HostConfig: &createContHostConfig}

	client.CreateContainer(createContOps)
	client.StartContainer("hello-docker", nil)

}

func main3() {
	endpoint := os.Getenv("DOCKER_HOST")
	path := os.Getenv("DOCKER_CERT_PATH")
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)
	client, _ := docker.NewTLSClient(endpoint, cert, key, ca)

	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}
}
