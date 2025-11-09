#! /bin/bash

set -e

REPO="dorei"

DIR="go-out"
mkdir -p "$DIR"

TARGETS=(
	"linux x86_64"
	"linux arm64"
	"darwin x86_64"
	"darwin arm64"
	"windows amd64"
	"windows arm64"
)

clear

for target in "${TARGETS[@]}"; do
	OS=$(echo $target | cut -d ' ' -f 1)
	ARCH=$(echo $target | cut -d ' ' -f 2)
	PROG="$REPO-$OS-$ARCH"
	OUTPUT="$DIR/$PROG"

	echo "Making $OS-$ARCH..."
	GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" -o "$OUTPUT" ./...
	
	if [ $? -ne 0 ]; then
		echo -e "\033[31m✗\033[0m Could not make $PROG"
		sleep 1
	else
		echo -e "\033[32m✓\033[0m Made $OUTPUT"
		sleep 1
	fi
done

echo

OS=$(go env GOOS)
ARCH=$(go env GOARCH)
PROG="$REPO-$OS-$ARCH"
TARGET_PATH="$HOME/.local/bin"

echo "Making $REPO..."
GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" -o "$REPO" ./...

if [ $? -ne 0 ]; then
	echo -e "\033[31m✗\033[0m Could not make $PROG"
	# sleep 1
else
	echo -e "\033[32m✓\033[0m Made $PROG"
	# sleep 1
fi
	
echo "Copying $REPO to $TARGET_PATH"
sudo cp "$REPO" "$TARGET_PATH/$REPO"

if [ $? -ne 0 ]; then
	echo
	echo -e "\033[31m✗\033[0m Could not copy $REPO to $TARGET_PATH"
	# sleep 1
else
	echo
	echo -e "\033[32m✓\033[0m Copied $REPO to $TARGET_PATH"
	# sleep 1
fi
