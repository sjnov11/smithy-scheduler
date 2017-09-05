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
    }
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
