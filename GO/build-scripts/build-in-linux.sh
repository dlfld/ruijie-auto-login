#! /bin/bash

TAG='1.3.1'
case "$1" in
    -t|--tag) TAG=$2;;
esac

dir_cur=`dirname $0`
cd "$dir_cur/.."

root_path=`pwd`
dir_out="$root_path/output"
path_src="$root_path/src/main.go"
path_config="./configuration.yaml"

lst_os=("darwin" "freebsd" "linux" "windows")
lst_arch=("386" "amd64" "arm")

# go env -w GO111MODULE=off

# build for all system with each architecture
cd $dir_out
CGO_ENABLED=0 
for os in ${lst_os[@]}; do
    for arch in ${lst_arch[@]}; do
        zip_base_name="ruijie-auto-login.${TAG}.${os}_${arch}"
        # output_filename="${dir_out}/${zip_base_name}"
        
        # append '.exe' at filename
        if [ $os == "windows" ]; then
            output_file="${zip_base_name}.exe"
        else
            output_file="${zip_base_name}"
        fi
        env GOOS=${os} GOARCH=${arch} go build -o ${output_file} ${path_src}
        
        # build failed
        if [ -f ${output_file} ]; then
            zip ${zip_base_name}.zip ${path_config} ${output_file}
            rm ${output_file}
        fi
        # echo "go build -o ${output_file} ${path_src}"
    done
done
 