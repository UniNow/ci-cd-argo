## Create and install repositories 

## Generate new deployment key

````shell
ssh-keygen -t ed25519 -C "demo-deploy@uninow.io"
````

Copy the public key for github
````shell
pbcopy < ~/.ssh/id_demo.pub
````

And the private key for the argo repositories
`````shell
pbcopy < ~/.ssh/id_demo
`````