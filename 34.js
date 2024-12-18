/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function(nums, target) {
    let l = 0
    let r = nums.length -1

    let lb = -1
    let rb = -1

    while(l <= r) {
        let m = Math.floor((l+r)/2)

        if (nums[m] === target && nums[m-1] !== target) {
            lb = m
        }

        if (nums[m] < target) {
            l = m + 1
        } else {
            r = m - 1
        }
    }

    l = 0
    r = nums.length -1
    while(l <= r) {
        let m = Math.floor((l+r)/2)

        if (nums[m] === target && nums[m+1] !== target) {
            rb = m
        }

        if (nums[m] <= target) {
            l = m + 1
        } else {
            r = m -1
        }
    }

    return [lb, rb]
};