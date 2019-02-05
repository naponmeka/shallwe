for GOOS in darwin linux
do
    for GOARCH in amd64
    do
        env GOOS=$GOOS GOARCH=$GOARCH go build -o ./bin/$GOOS-$GOARCH/shallwe
    done
done


# env GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/shallwe