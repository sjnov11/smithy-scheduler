var timeTable = createTimeTablePainter();

function createTimeTablePainter() {
  return new Vue({
    el: '#timeTable',
    data: {
      startHour : 9,
      endHour : 20,

      lectureDataArray : null,

      rowNames: [], // time
      colNames: [], // day
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
      // 0 is sunday. 6 is saturday.
      var day = 0;
      do {
        this.colNames.push(day);
        day++;
      } while (day <= 6);
    },

    methods: {
      drawTimeTable: function() {
        if (this.lectureDataArray == null) {
          console.log("(drawTimeTable) failed. There is no lecture data")
          return false;
        }

        this.lectureDataArray.forEach(function(lecture) {
          var Start_time;
          var Start_minute;
          var End_time;
          var End_minute;



          var timeTable = $("#timeTable");
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


            var first_cell = timeTable.find("div[hour='"+Start_time+"'][minute='"+Start_minute+"']").find("div[day='"+Day+"']");
            first_cell[0].innerHTML = lecture.GwamokNm;

            var count = 0;
            do {
              var cell = timeTable.find("div[hour='"+Start_time+"'][minute='"+Start_minute+"']").find("div[day='"+Day+"']");
              cell.css('background', '#cee');

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


        })
      },
    }
  });
};

timeTable.lectureDataArray = possibleData;
timeTable.drawTimeTable();
