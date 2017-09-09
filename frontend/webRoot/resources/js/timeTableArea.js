var timeTableArea = new Vue({
  el: '#timeTableArea',
  data: {
    startHour : 9,
    endHour : 20,

    timeTableData: null, // array of each time table data

    rowNames: [], // time
    colNames: [], // day

    firstRow: ["","일","월","화","수","목","금","토"], // day name
    firstCol: [], // time name

    startButtonOfBasket : $('.combination-start.button'),
  },

  created: function() {
    //init rowNames
    var hh = this.startHour;
    do {
      var timeStructure = {
        hour : null,
        minute : null
      };

      var hhStr = hh.toString();
      if (hh < 10) {
        hhStr = "0" + hhStr;
      }

      timeStructure.hour = hhStr;
      timeStructure.minute = "00";

      this.rowNames.push(timeStructure);

      timeStructure = {
        hour : null,
        minute : null
      };

      timeStructure.hour = hhStr;
      timeStructure.minute = "30";

      this.rowNames.push(timeStructure);

      hh++;
    } while (hh <= this.endHour);

    //init colNames
    // -1 is Name, 0 is sunday. 6 is saturday.
    var day = 0;
    do {
      this.colNames.push(day);
      day++;
    } while (day <= 6);
  },

  methods: {
    setTimeTableData : function(data) {
      // time table space is generated when the timeTableData is changed.
      this.timeTableData = data;

      var vueObject = this;
      // wait 500ms to html code changing.
      sleep(500).then(() => {
        vueObject.drawTimeTable();

        // restore button
        this.startButtonOfBasket.popup({
          popup : $(".combination-complete.popup"),
          position : "top center",
          on    : 'click'
        });

        if (data.length != 0) {
          this.startButtonOfBasket.popup('show').removeClass('loading');
        }

        sleep(2000).then(() => {
          this.startButtonOfBasket.popup('hide');
          this.startButtonOfBasket.popup('destroy');
        })

      });
    },
    drawTimeTable: function() {

      for (var tableNumber = 0, tableLength = this.timeTableData.length; tableNumber < tableLength; tableNumber++) {
        var lectureData = this.timeTableData[tableNumber];

        if (lectureData == null) {
          console.log("(drawTimeTable) failed. There is no lecture data")
          return false;
        }

        for (var lectureIdx = 0, l = lectureData.length; lectureIdx < l; lectureIdx++) {


          var lecture = lectureData[lectureIdx];

          var Start_time;
          var Start_minute;
          var End_time;
          var End_minute;


          var timeTable = $(this.$el);
          lecture.SecondData.TimesAndClass.forEach(function(timeAndClass) {
            Start_time   = timeAndClass.Start_time;
            Start_minute = timeAndClass.Start_minute;
            End_time   = timeAndClass.End_time;
            End_minute = timeAndClass.End_minute;
            Day          = timeAndClass.Day;

            switch (Day) {
              case '일': Day = 0; break;
              case '월': Day = 1; break;
              case '화': Day = 2; break;
              case '수': Day = 3; break;
              case '목': Day = 4; break;
              case '금': Day = 5; break;
              case '토': Day = 6; break;
              default:
                console.log("no day data.");
                return;
            }


            // lecture name display
            var first_cell = timeTable.find("div.time-table[table-number='"+tableNumber.toString()+"']").find("div[hour='"+Start_time+"'][minute='"+Start_minute+"']").find("div[day='"+Day+"']");
            first_cell[0].innerHTML = lecture.GwamokNm;

            var count = 0;
            do {
              // coloring the background
              var cell = timeTable.find("div.time-table[table-number='"+tableNumber.toString()+"']").find("div[hour='"+Start_time+"'][minute='"+Start_minute+"']").find("div[day='"+Day+"']");
              cell.css('background', '#e8e8e8');
              cell.css('border', '1px solid #e8e8e8');

              // Increase Strat_time
              if (Start_minute == "30") {
                Start_minute = "00";

                Start_time   = parseInt(Start_time);
                Start_time++;
                if (Start_time < 10) {
                  Start_time = "0" + Start_time.toString();
                } else {
                  Start_time = Start_time.toString();
                }
              } else {
                Start_minute = "30";
              }

              var truth_value = !(Start_time == End_time && Start_minute == End_minute);

              // prevent infinite loop
              count++
              if (count > 1000) {
                console.log("infinite loop has been occurred.");
                console.log(lecture.GwamokNm);
                break;
              }
            } while ( truth_value );
          });
        }
      }
    },
  },
  updated: function() {
    // modal is closed when a time table is clicked.

    if (this.timeTableData != null && this.timeTableData.length != 0) {
      $('.combination-result.modal')
        .modal('attach events', '.time-table', 'hide')
      ;
    }

    var numberOfTables = this.timeTableData.length;
    var thumbnailText = document.getElementById("howManyTableIsGenerated");
    thumbnailText.innerHTML = "Possible<br/>Time Table<br/>Count:" + numberOfTables.toString();
    thumbnailText.style.margin = "44px 0 0 0";
  }
});
