var combinationConditionDimmer = new Vue({
  el : '#combinationConditionDimmer',
  data : {
    // minimum range
    // if minLectureNumber is 2 and minLectureNumberRange is 3, displayed number is 2,3,4
    minLectureNumber : 2,
    minLectureNumberRange : 3,

    // maximum range
    // if maxLectureNumber is 10 and maxLectureNumberRange is 3, displayed number is 8,9,10
    maxLectureNumber : 10,
    maxLectureNumberRagne : 3,
  },
});
