


$(document).ready(function(){
	var $boxes = $('.box');
	var $spheres = $('.sphere');
	var $boxcontainer = $('.boxcontainer');
	var $paragraph = $('p');
	var $spherecontainer = $('.spherecontainer');
	
	/*Hide the boxes, then fade in individually on load */
	$boxes.hide();
	$boxes.each(function(index){
		var boxcolor = '#'+Math.random().toString(16).substr(-6);
		$(this).css({
			'background-color': boxcolor,
		});
    });
	
	/*Animate the title of the page to grow in size on load*/
	$paragraph.hide();
    setTimeout(function(){
		$paragraph.show();
		$paragraph.animate({fontSize: "200%" }, 1000 );
	},7000);
	
	/*slide or fade boxes into the page*/
	$boxes.each(function(index){
		$(this).delay(170*index).fadeIn(300);
	});
	
	/*slide or fade spherecontainer into view*/
		$spherecontainer.hide();
		$spherecontainer.delay(8000).fadeIn(1000);

	/*This is the on spheres click function, which removes spheres*/
	$spheres.on('dblclick', function(){
		$(this).remove();
	});
	
	
	/*This is the on boxes click function, which removes boxes and presents easter egg if no boxes are visible*/
	$boxes.on('dblclick', function(){
		$(this).remove();
		numboxes = $('.box:visible').length;
		if(numboxes < 1){
			$boxcontainer.css({
				'background-image': 'url(\"' + 'easter-egg.jpg' + '\")',
		        'background-repeat': 'no-repeat',	
				'background-position': 'center',
			});
		};
	});	
	
	
	
});