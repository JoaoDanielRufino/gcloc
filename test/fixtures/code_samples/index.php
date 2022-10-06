<?php

// pass by value
function valFun($num) {
	$num += 2;
	return $num;
}

// pass by reference
function refFun(&$num) {
	$num += 10;
	return $num;
}

$n = 10;

valFun($n);
echo "The original value is still $n \n";

refFun($n);
echo "The original value changes to $n";

?>
