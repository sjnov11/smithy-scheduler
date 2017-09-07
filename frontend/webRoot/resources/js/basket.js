function showBasket() {
  $("#basket").css("display","");
  $("#searchForm").addClass("subject-selected");
};

var selectedLectureBasketList = new Vue({
  el: '#selectedLectureList',
  data: {
    lectureData: []
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
