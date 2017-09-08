function showBasket() {
  $("#basket").css("display","");
  $("#searchForm").addClass("subject-selected");
};

var selectedLectureBasketList = new Vue({
  el: '#selectedLectureList',
  data: {
    lectureData: [],
    checkedLectureData: [],
    timeTablesData: []
  },
  mounted: function() {
    combiStartButton = $(this.$el).find('.combination-start.button').popup({
      popup : $(".combination-complete.popup"),
      position : "top center",
      on    : 'click'
    });

    // use closure
    var vueObject = this;
    function generateTimeTablesData() {
      vueObject.checkedLectureData = [];

      // remove lecture that does not have 'checked'
      var lectures = $(vueObject.$el).find('.each-selected-lecture');
      for (var i = 0, l = lectures.length; i < l; i++) {
        var lecture = lectures[i];
        if( $(lecture).find(".ui.checkbox.checked").length != 0) {
          vueObject.checkedLectureData.push(vueObject.lectureData[i]);
        }
      }
      vueObject.timeTablesData = vueObject.generatePossibleTimeTableData();
      console.log(vueObject.timeTablesData);
    }
    combiStartButton.click(generateTimeTablesData);
  },
  methods: {
    add: function(SuupNo2, BanSosokNm){
      // majorDataMap is a global variable in subjectTable.js
      var lectures = majorDataMap.get(BanSosokNm);

      // find the lecture which is same with SuupNo2
      var idx = binarySearchForLectureArray.call(lectures, SuupNo2);
      if (idx == -1) {
        console.log("(add, binarySearch..) No lecture data. Add is failed.");
        return
      }
      this.lectureData.push(lectures[idx]);
    },

    remove: function(SuupNo2) {
      var idx = -1;

      for (var i = 0, l = this.lectureData.length; i < l; i++) {
        if(SuupNo2 == this.lectureData[i].SuupNo2) {
          idx = i;
          break;
        }
      }

      if (idx == -1) {
        console.log("(remove, binarySearch..) No lecture data. Remove is failed.");
        return
      }

      for (var i = idx, l = this.lectureData.length; i < l - 1; i++) {
        this.lectureData[i] = this.lectureData[i+1];
      }
      this.lectureData.pop();
    },

    moveUp: function(SuupNo2) {
      if (this.lectureData.length == 0) {
        // no data
        return;
      }


      var idx;

      for (var i = 0, l = this.lectureData.length; i < l; i++) {
        if(SuupNo2 == this.lectureData[i].SuupNo2) {
          idx = i;
          break;
        }
      }

      // when lecture is on top.
      if (idx == 0) {
        return;
      }

      var tempLecture = this.lectureData[idx];
      this.lectureData[idx] = this.lectureData[idx-1];
      this.lectureData[idx-1] = tempLecture;
      this.$forceUpdate();
    },

    moveDown: function(SuupNo2) {
      if (this.lectureData.length == 0) {
        // no data
        return;
      }

      var idx;

      for (var i = 0, l = this.lectureData.length; i < l; i++) {
        if(SuupNo2 == this.lectureData[i].SuupNo2) {
          idx = i;
          break;
        }
      }

      // when the lecture is last one or does not exist
      if (idx >= this.lectureData.length - 1) {
        return;
      }

      var tempLecture = this.lectureData[idx];
      this.lectureData[idx] = this.lectureData[idx+1];
      this.lectureData[idx+1] = tempLecture;
      this.$forceUpdate();
    },
    isPossibleData: function(lectureDataArray) {
      // 2차원 맵으로 시간표의 중복을 체크한다.

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


      var timeTable = new Map();

      for (var idx in lectureDataArray) {
        var lectureData = lectureDataArray[idx]

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
            if (timeTable.get(start) == undefined) {
              timeTable.set(start, new Map());
            }

            // check whether time is duplicated
            var duplicated = timeTable.get(start).get(day);
            if (duplicated) {
              return false;
            } else {
              timeTable.get(start).set(day, true);
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
    generatePossibleTimeTableData: function(minLectureNumber, maxLectureNumber) {
      if (this.checkedLectureData.length == 0) {
        console.log("there is no checked lecture");
        return
      }

      // This powerSet function is from 
      // https://codereview.stackexchange.com/questions/7001/generating-all-combinations-of-an-array
      var powerSet = function (list, minLength, maxLength){
        var set = [],
          listSize = list.length,
          combinationsCount = (1 << listSize),
          combination;

        for (var i = 1; i < combinationsCount ; i++ ){
          var combination = [];

          // check the number of 1s.
          var bits = i;
          var numberOfElements = 0;
          do {
            if (bits & 1 == 1) {
              numberOfElements++;
            }
            bits >>= 1;
          } while (bits > 0);

          if (numberOfElements < minLength || numberOfElements > maxLength) {
            // it does not meet the requirements.
            continue;
          }

          for (var j=0;j<listSize;j++){
            if ((i & (1 << j))){
              combination.push(list[j]);
            }
          }
          set.push(combination);
        }
        return set;
      };

      var allSets = powerSet(this.checkedLectureData, minLectureNumber, maxLectureNumber);

      var resultArray = [];

      for (var idx in allSets) {
        if (this.isPossibleData(allSets[idx])) {
          resultArray.push(allSets[idx]);
        }
      }

      return resultArray;
    }
  },

  updated: function(){
    // make last added element checked
    var checkbox = $(this.$el).find('.ui.checkbox:last').checkbox('check');

    // add tooltip for each lecture
    var last_added_lecture = $(this.$el).find('.each-selected-lecture:last');

    last_added_lecture.popup({
      popup: last_added_lecture.find('.each-lecture.popup'),
      position: "left center",
    })
    ;

    // checkbox events
    var checkboxClick = function() {
      if ($(this).hasClass('checked')) {
        $(this).parent().css('background-color', '#f4ffdb');
      } else {
        $(this).parent().css('background-color', '#f2f2f2');
      }
    };


    var checkboxMouseenter = function() {
      if ($(this).hasClass('checked')) {
        $(this).parent().css('background-color', '#f4ffdb');
      } else {
        $(this).parent().css('background-color', '#f2f2f2');
      }
    };

    var checkboxMouseleave = function(){
      $(this).parent().css('background-color', '');
    };


    var checkboxMousedown = function() {
      if ($(this).hasClass('checked')) {
        $(this).parent().css('background-color', '#ebffbf');
      } else {
        $(this).parent().css('background-color', '#e4e4e4');
      }
    };
    var checkboxMouseup = function() {
      if ($(this).hasClass('checked')) {
        $(this).parent().css('background-color', '#f4ffdb');
      } else {
        $(this).parent().css('background-color', '#f2f2f2');
      }
    };


    // checkbox binding.
    checkbox.unbind("click", "checkboxClick");
    checkbox.unbind("mouseenter", "checkboxMouseenter");
    checkbox.unbind("mouseleave", "checkboxMouseleave");
    checkbox.unbind("mousedown", "checkboxMousedown");
    checkbox.unbind("mouseup", "checkboxMouseup");

    checkbox.bind({
      click: checkboxClick,
      mouseenter: checkboxMouseenter,
      mouseleave: checkboxMouseleave,
      mousedown: checkboxMousedown,
      mouseup: checkboxMouseup
    });

    // add event to each button
    var last_buttons = $(this.$el).find('.each-lecture-buttons:last');

    // remove event lister and add.
    var left_button = last_buttons.find('.left.button');
    left_button.unbind("click");
    left_button.unbind("mouseenter");
    left_button.unbind("mouseleave");

    //bind
    left_button.bind({
      click: function() {
        var number = $(this).parent().parent().attr('lecture-number');
        var name = $(this).parent().parent().attr('lecture-name');


        selectedLecture.delete(number);

        //fucking dirty code... shit
        var modal_lecture = $("div.ui.dimmer.modals.page.transition.hidden").find(".each-subject.modal-subject-name[lecture-number='"+ number  +"']");
        modal_lecture.removeClass("selected");


        sameNameLectures = modal_lecture.parent().find(".selected");
        if (sameNameLectures.length == 0){
          $("td.selectable[lecture-name='" + name + "']").removeClass("selected");
        }


        // Vue
        selectedLectureBasketList.remove(number);
      },

      mouseenter: function() {
        $(this).parent().parent().css('background-color', '#fff0f0');
      },
      mouseleave: function(){
        $(this).parent().parent().css('background-color', '');
      },
      mousedown: function(){
        $(this).parent().parent().css('background-color', '#ffdada');
      },
      mouseup: function(){
        $(this).parent().parent().css('background-color', '#fff0f0');
      }
    });

    // up button
    var right_top_button = last_buttons.find('.right.top.button');
    right_top_button.unbind("click");
    right_top_button.unbind("mouseenter");
    right_top_button.unbind("mouseleave");

    right_top_button.bind({
      click: function() {
        var n = $(this).parent().parent().attr('lecture-number');
        selectedLectureBasketList.moveUp(n);
      },

      mouseenter: function() {
        $(this).parent().parent().css('background-color', '#ffffea');
      },
      mouseleave: function(){
        $(this).parent().parent().css('background-color', '');
      },
      mousedown: function(){
        $(this).parent().parent().css('background-color', '#ffffc0');
      },
      mouseup: function(){
        $(this).parent().parent().css('background-color', '#ffffea');
      }
    });


    // down button
    var right_bottom_button = last_buttons.find('.right.bottom.button');
    right_bottom_button.unbind("click");
    right_bottom_button.unbind("mouseenter");
    right_bottom_button.unbind("mouseleave");

    right_bottom_button.bind({
      click: function() {
        var n = $(this).parent().parent().attr('lecture-number');
        selectedLectureBasketList.moveDown(n);
      },
      mouseenter: function() {
        $(this).parent().parent().css('background-color', '#f0ffff');
      },
      mouseleave: function(){
        $(this).parent().parent().css('background-color', '');
      },
      mousedown: function(){
        $(this).parent().parent().css('background-color', '#daffff');
      },
      mouseup: function(){
        $(this).parent().parent().css('background-color', '#f0ffff');
      }
    });
  }
});


function updateSelectedLectureList(){
  console.log(selectedLecture); 

  // compare selectedLectureBasketList and selectedLecture to find which lecture is needed to update

  for (var key of selectedLecture.keys()) {
    console.log(key);
  }

  // add more remove lecture in selectedLectureBasketList
}



/**
 * Performs a binary search on the host array. This method can either be
 * injected into Array.prototype or called with a specified scope like this:
 * binaryIndexOf.call(someArray, searchElement);
 *
 * @param {*} searchElement The item to search for within the array.
 * @return {Number} The index of the element which defaults to -1 when not found.
 */

// https://oli.me.uk/2013/06/08/searching-javascript-arrays-with-a-binary-search/
function binarySearchForLectureArray(searchElement) {
  'use strict';

  var minIndex = 0;
  var maxIndex = this.length - 1;
  var currentIndex;
  var currentElement;

  while (minIndex <= maxIndex) {
    currentIndex = (minIndex + maxIndex) / 2 | 0;
    currentElement = this[currentIndex].SuupNo2;

    if (currentElement < searchElement) {
      minIndex = currentIndex + 1;
    }
    else if (currentElement > searchElement) {
      maxIndex = currentIndex - 1;
    }
    else {
      return currentIndex;
    }
  }

  return -1;
}
