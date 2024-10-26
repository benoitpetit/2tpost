#!/bin/bash

# Define paths for builds
LINUX_BUILD_DIR="build/bin/linux"
WINDOWS_BUILD_DIR="build/bin/windows"
MAC_BUILD_DIR="build/bin/mac"

# Create build directories if necessary
mkdir -p $LINUX_BUILD_DIR $WINDOWS_BUILD_DIR $MAC_BUILD_DIR

# Build for Linux with the -tags webkit2_41 option
echo "Building for Linux..."
wails build -clean -tags webkit2_41 -platform linux/amd64 -o $LINUX_BUILD_DIR/2tpost

# Check if the Linux build succeeded
if [ $? -eq 0 ]; then
    echo "Linux build successful!"
else
    echo "Linux build failed!"
    exit 1
fi

# Build for Windows with the -obfuscated option
echo "Building for Windows..."
wails build -platform windows/amd64 -o $WINDOWS_BUILD_DIR/2tpost.exe

# Check if the Windows build succeeded
if [ $? -eq 0 ]; then
    echo "Windows build successful!"
else
    echo "Windows build failed!"
    exit 1
fi

# Build for macOS
echo "Building for macOS..."
wails build -platform darwin/amd64 -o $MAC_BUILD_DIR/2tpost

# Check if the macOS build succeeded
if [ $? -eq 0 ]; then
    echo "macOS build successful!"
else
    echo "macOS build failed!"
    exit 1
fi

echo "All builds completed successfully!"
