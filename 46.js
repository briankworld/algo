function permute(nums) {
    const result = []

    function bt(index) {
        if (index === nums.length - 1) {
            result.push([...nums])
            return
        }

        for(let i = index; i < nums.length; i ++) {
            [nums[index], nums[i]] = [nums[i], nums[index]]
            bt(index + 1);
            [nums[index], nums[i]] = [nums[i], nums[index]]
        }
    }

    bt(0)
    return result
}