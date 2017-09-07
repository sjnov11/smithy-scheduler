var timeTable = new Vue({
  el: '#timeTable',
  data: {
    tableData : null,
    lectureDataArray : null,

    rowNames: [], // time
    colNames: [], // day
  },
  created: function() {
    //init rowNames
    var time = 900;
    do {
      this.rowNames.push(time);

      if (time % 100 == 0) {
        time += 30;
      } else {
        time += 70;
      }

    } while (time <= 2030);
    
    //init colNames

    // 0 is sunday. 6 is saturday.
    var day = 0;
    do {
      this.colNames.push(day);
      day++;
    } while (day <= 6);
  },

  methods: {
    generateTableData: function() {
      // 2차원 맵으로 시간표 데이터를 만든다.

      //        일(0)  월(1)  화(2)  수(3)  목(4)  금(5)  토(6)
      // 0700
      // 0730
      // 0800
      // 0830
      // 0900
      // 0930
      // 1000
      // 1030
      // .
      // .
      // .


      this.tableData = new Map();

      for (var idx in this.lectureDataArray) {
        var lectureData = this.lectureDataArray[idx]

        if (lectureData.SecondData.TimesAndClass.length == 0) {
          // console.log("no SecondData. continue for loop");
          continue;
        }

        // each lecture, iterate all time
        for (var j in lectureData.SecondData.TimesAndClass) {
          var classTimeData = lectureData.SecondData.TimesAndClass[j]

          // below three variables are all int.
          var day;
          var start;
          var end;

          switch (classTimeData.Day) {
            case '일': day = 0; break;
            case '월': day = 1; break;
            case '화': day = 2; break;
            case '수': day = 3; break;
            case '목': day = 4; break;
            case '금': day = 5; break;
            case '토': day = 6; break;
            default: 
              console.log("no time data. continue for loop");
              continue;
          }
          // console.log(day);
          start = parseInt(classTimeData.Start_time + classTimeData.Start_minute);

          end = parseInt(classTimeData.End_time + classTimeData.End_minute);

          // 30으로 끝나면 70을 더하고, 00으로 끝나면 30을 더해서 계속 체크해가자.

          do {
            // row는 시간, col이 요일이다.
            if (this.tableData.get(start) == undefined) {
              this.tableData.set(start, new Map());
            }

            // check whether time is duplicated
            var duplicated = this.tableData.get(start).get(day);
            if (duplicated) {
              this.tableData = null;
              return false;
            } else {
              // idx is the index of this.lectureDataArray
              this.tableData.get(start).set(day, idx);
            }

            // check next time
            if (start % 100 == 0) {
              start += 30;
            } else {
              start += 70;
            }

          } while (start != end);

        }

      }
      return true;
    },
    getLectureIdx: function(time, day) {
      return this.tableData.get(time).get(day)
    }
  }
});


// No support map sturcture in vue.
// How...?
timeTable.lectureDataArray = possibleData;
timeTable.generateTableData();
