#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

echo "Building Docker image microservice1"
docker build -t microservice1 -f Dockerfile .

echo "Generating ./run.sh"
rm -f ./run.sh
echo "#!/bin/bash" >> ./run.sh
echo "docker run --net=codefresh_test -p 3000:3000 --rm --name ms1 -it microservice1" >> ./run.sh
chmod +x ./run.sh

echo "You can now run ./run.sh to start microservice1."



