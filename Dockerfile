FROM busybox
ADD views /gopkg/src/github.com/qor/qor/admin/views
ADD i18nviews /gopkg/src/github.com/qor/qor/i18n/views
ENV GOPATH /gopkg
ADD devicem /bin/
CMD /bin/devicem
