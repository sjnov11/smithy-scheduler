function sleep (time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}


var majorIsSelected = true;

$('.ui.dropdown')
  .dropdown()
;

$('#btn_login').click(function() {
  $('.login.page').dimmer('show');
});

$('#cCondition').click(function() {
  $('.config-c-condition.page').dimmer('show');
});


$('#btn_major').click(function() {
  $(this).removeClass("basic");
  $('#btn_cultural').addClass("basic");

  $('#majorSearchForm').css('display','');
  $('#culturalSearchForm').css('display', 'none');

  majorIsSelected = true;
  displayMajorTable();


  // Update input value
  var formText = $("input[name='receipt_major']").val();
  $('#echoInputValue').text(formText)
});

$('#btn_cultural').click(function() {
  $(this).removeClass("basic");
  $('#btn_major').addClass("basic");

  $('#culturalSearchForm').css('display','');
  $('#majorSearchForm').css('display', 'none');

  majorIsSelected = false;
  hideMajorTable();

  // Update input value
  var formText = $("input[name='receipt_cultural']").val();
  $('#echoInputValue').text(formText)
});


$('#btn_major_selected').click(function() {
  location.href="major-selected.html"
});

$('button.lecture-select').click(function() {
  var cell = $(this).parent().parent().parent().parent();

  if(cell.hasClass('positive')) {
    cell.removeClass('positive');
  } else {
    cell.addClass('positive');
  }
});

$('#timeTableThumbnail').click(function() {
  // don't show result when generated table is zero.
  if (selectedLectureBasketList.checkedLectureData.length != 0) {
    $('.ui.page.dimmer.combination-result').dimmer('show');
  }
});

$('#btn_request_post').click(function(event) {
  event.preventDefault();

  var x = new XMLHttpRequest();

  x.open("POST", "/db/");
  x.onreadystatechange = function() {
    if(x.readyState === 4 && x.status === 200) {
      // console.log(x.response);
      alert(x.response)
    }
  }

  x.send();
});
