var timeTableArea = new Vue({
  el: '#timeTableArea',
  data: {
    timeTableData: [], // array of each time table data
  },

  methods: {
    setTimeTableData : function(data) {
      // time table space is generated when the timeTableData is changed.
      this.timeTableData = data;
    },
  },
  updated: function() {
    // modal is closed when a time table is clicked.
    $('.combination-result.modal')
      .modal('attach events', '.time-table', 'hide')
    ;
  }
});
