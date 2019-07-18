export GOPATH=/ws:/ws/jstack-thirdparty-libs;

pushd ${SRC} &&
count=$(git rev-list --count ${COMMIT_SHA}) &&
iteration=${ADDITIONAL_VERSION}.$(git rev-parse --short ${COMMIT_SHA}) &&
BUILD_TIME=$(git log -1 --format=%ci | sed 's/ /T/1;s/ //g') &&
env GOOS=linux GOARCH=amd64 GO_GCFLAGS=-N go build -ldflags "-X jd.com/jstack-common/tag.GitVersion=${MAJOR_VERSION}.${count}-$iteration-$BRANCH -X jd.com/jstack-common/tag.BuildTime=$BUILD_TIME" -o ${NAME} -a api/main.go;
env GOOS=linux GOARCH=amd64 GO_GCFLAGS=-N go build -ldflags "-X jd.com/jstack-common/tag.GitVersion=${MAJOR_VERSION}.${count}-$iteration-$BRANCH -X jd.com/jstack-common/tag.BuildTime=$BUILD_TIME" -o ${CLI_NAME} -a cli/main.go;

go run cli/bash_auto_complete/generator.go ;
popd;

mkdir -p rpm_build/usr/local/bin && mkdir -p rpm_build/usr/local/share/${NAME}/samples && mkdir -p rpm_build/etc/bash_completion.d &&
mv ${SRC}/${NAME} rpm_build/usr/local/bin &&
mv ${SRC}/${CLI_NAME} rpm_build/usr/local/bin &&
mv ${SRC}/ccs.bash_complete rpm_build/etc/bash_completion.d &&
cp ${SRC}/api/etc/*.json rpm_build/usr/local/share/${NAME}/samples &&


cd rpm_build && \
fpm -f -s dir -t rpm -n cc-server -v $MAJOR_VERSION.${count} \
--iteration $iteration --package .. -a x86_64 . ;