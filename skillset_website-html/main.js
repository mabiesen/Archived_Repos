function main(){
$('.skillset').hide();
  $('.skillset').fadeIn(1000);
$('.projects').hide();
  $('.projects-button').on('click',function(){
    $(this).next().toggle();
    $(this).toggleClass('active');
  });
  $('.learning-button').on('click',function(){
    $(this).next().toggle();
    $(this).toggleClass('active');
  });
 }

$(document).ready(main);