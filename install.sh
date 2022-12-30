#!/usr/bin/env bash

echo "Installing GoJelastic..."
echo "------------------------"

# Determining the Linux distribution and architecture
distro=$(lsb_release -i -s)
arch=$(uname -m)

echo "Distribution: $distro"
echo "Architecture: $arch"

# GoJelastic version
version="v0.1.3"

echo "Version: $version"
echo "------------------------"

# URL for downloading the archive based on the distribution and architecture
url=""

case "$distro" in
  "Darwin")
    case "$arch" in
      "x86_64")
        url="https://github.com/yoanbernabeu/GoJelastic/releases/download/${version}/GoJelastic-${version}-darwin-amd64.tar.gz"
        ;;
      "arm64")
        url="https://github.com/yoanbernabeu/GoJelastic/releases/download/${version}/GoJelastic-${version}-darwin-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  "Ubuntu"|"Debian"|"Raspbian")
  echo "Downloading GoJelastic..."
    case "$arch" in
      "i686")
        url="https://github.com/yoanbernabeu/GoJelastic/releases/download/${version}/GoJelastic-${version}-linux-386.tar.gz"
        ;;
      "x86_64")
        url="https://github.com/yoanbernabeu/GoJelastic/releases/download/${version}/GoJelastic-${version}-linux-amd64.tar.gz"
        echo $url
        ;;
      "arm64")
        url="https://github.com/yoanbernabeu/GoJelastic/releases/download/${version}/GoJelastic-${version}-linux-arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Unsupported distribution"
    exit 1
    ;;
esac

# Downloading the archive to home directory (and check if url is not 404)
echo "Downloading GoJelastic..."
wget -q --spider $url
if [ $? -eq 0 ]; then
  wget -O ~/GoJelastic.tar.gz $url -q --show-progress
else
  echo "------------------------"
  echo "GoJelastic archive not found"
  echo "------------------------"
  exit 1
fi

# Extracting the archive (if it exists)
echo "Extracting GoJelastic..."
if [ -f ~/GoJelastic.tar.gz ]; then
  tar -xzf ~/GoJelastic.tar.gz -C ~/
else
  echo "GoJelastic archive not found"
  exit 1
fi

# Removing the archive
echo "Removing archive..."
rm ~/GoJelastic.tar.gz

# Moving the binary to /usr/local/bin
echo "Moving GoJelastic to /usr/local/bin..."
sudo mv ~/GoJelastic /usr/local/bin/

# Making the binary executable
echo "Making GoJelastic executable..."
sudo chmod +x /usr/local/bin/GoJelastic

# Sending a message to the user
echo "-----------------------------------------"
echo "GoJelastic successfully installed"
echo "-----------------------------------------"