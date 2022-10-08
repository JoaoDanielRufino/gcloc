=begin
Given an array of 0s and 1s in random order.
Segregate 0s on left side and 1s on right side of the array
=end
def segregate(a)
  left, right = 0, a.length - 1
  while left < right
    # Increment left index while we see 0 at left
    while a[left] == 0
      left += 1
    end

    # Decrement right index while we see 1 at right
    while a[right] == 1
      right -= 1
    end

    # Exchange arr[left] and arr[right]
    if left < right
      a[left], a[right] = a[right], a[left]
      left += 1
      right -= 1
    end
  end

  return a
end

p segregate([1,0,1,1,0,0])
# Output: [0, 0, 0, 1, 1, 1]
