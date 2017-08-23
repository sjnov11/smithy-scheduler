new Vue({
  el: '#majorList',
  data: {
    majors: [
    {{range .}} {{ .MajorName }}
    {{end}}
    ]
  }
});
