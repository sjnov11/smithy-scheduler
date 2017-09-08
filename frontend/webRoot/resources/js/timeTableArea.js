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

    var numberOfTables = this.timeTableData.length;
    var thumbnailText = document.getElementById("howManyTableIsGenerated");
    thumbnailText.innerHTML = "Possible<br/>Time Table<br/>Count:" + numberOfTables.toString();
    thumbnailText.style.margin = "44px 0 0 0";
  }
});
