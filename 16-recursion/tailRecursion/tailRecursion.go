package tailRecursion

func BouncingBall(h, bounce, window float64) int {
  return tailBounce(h, bounce, window, -1)
}

func tailBounce(h, bounce, window float64, bounceCount int) int {
  if h <= window || bounce <= 0 || bounce >= 1 {
    return bounceCount
  }
  
  rebounce := h*bounce
  return tailBounce(rebounce, bounce, window, bounceCount + 2)
}

// solution to https://www.codewars.com/kata/5544c7a5cb454edb3c000047/train/go