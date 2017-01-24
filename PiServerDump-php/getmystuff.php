<?php include('simple_html_dom.php');

function get_gold(){
$html = file_get_html('https://www.kitco.com/charts/livegold.html');
foreach($html->find('span#sp-bid') as $e){
	return $e->innertext . '<br>';
	}
}


function get_silver(){
$html2 = file_get_html('https://www.kitco.com/charts/livesilver.html');
foreach($html2->find('span#sp-bid') as $e){
        return $e->innertext . '<br>';
        }
}

function clinton_electoral(){
$html = file_get_html('http://www.politico.com/2016-election/results/map/president');
foreach($html->find('.balance-stats .type-democrat .stats-content .stats-secondary .macro') as $e){
        return $e->innertext;
}
}

function donald_electoral(){
$html = file_get_html('http://www.politico.com/2016-election/results/map/president');
foreach($html->find('.balance-stats .type-republican .stats-content .stats-secondary .macro') as $e){
        return $e->innertext;
}
}


$getgold = get_gold();
$getsilver = get_silver();
$hillary = hillary_electoral();
$donald = donald_electoral();
?>

