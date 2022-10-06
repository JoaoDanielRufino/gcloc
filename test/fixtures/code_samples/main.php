<?php
// Recursive PHP program for insertion sort

// Recursive function to sort an
// array using insertion sort
function insertionSortRecursive(&$arr, $n)
{

    // Base case
    if ($n <= 1)
        return;

    // Sort first n-1 elements
    insertionSortRecursive($arr, $n - 1);

    // Insert last element at its correct
    // position in sorted array.
    $last = $arr[$n - 1];
    $j = $n - 2;

    // Move elements of arr[0..i-1], that are
    // greater than key, to one position ahead
    // of their current position
    while ($j >= 0 && $arr[$j] > $last)
    {
        $arr[$j + 1] = $arr[$j];
        $j--;
    }
    $arr[$j + 1] = $last;
}

// A utility function to
// print an array of size n
function printArray(&$arr, $n)
{
    for ($i = 0; $i < $n; $i++)
        echo $arr[$i]." ";
}

// Driver Code
$arr = array(12, 11, 13, 5, 6);
$n = sizeof($arr);

insertionSortRecursive($arr, $n);
printArray($arr, $n);

?>