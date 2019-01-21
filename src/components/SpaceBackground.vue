<template>
  <div class="space">
    <canvas id="globalBackground"></canvas>
  </div>
</template>

<script>

export default {
  mounted() {
    let canvas = document.getElementById("globalBackground")
    let context = canvas.getContext('2d')


    window.onload = function () {
      canvas.width = window.innerWidth
      canvas.height = window.innerHeight
    }

    function random(max, min) {
      min = min || 0
      
      return chance.integer({ min: min, max: max })
    }

    let stars = new Array(100).fill().map(function (element) {
      element = {}
      element.size = 1,
      element.x = random(canvas.offsetWidth)
      element.y = random(canvas.offsetHeight)
      element.color = "#FFFFFF"
      element.shadowColor='#0000FF';
      element.shadowBlur= 20;
      element.y = random(canvas.height)
      return element
    })
  
    function moveStars(arr) {
      for (let inx = 0; inx < arr.length; inx++) {
        const star = arr[inx];
        if (star.x <= 0 || star.x >= canvas.offsetWidth) {
          star.x = random(canvas.offsetWidth)
        } else  {
          if (star.x >= canvas.offsetWidth/2 ) {
            star.x += 1
          } else {
            star.x -= 1
          }
        }
        if (star.y <= 0 || star.y >= canvas.offsetHeight) {
          star.y = random(canvas.offsetHeight)
        } else {
          if (star.y >= canvas.offsetHeight/2) {
            star.y += 1
          } else {
            star.y -= 1
          }
        }
      }
    }

    let positionChangeFlag = true
    function draw() {
      context.fillStyle = "#000000"
      context.fillRect(0, 0, canvas.width, canvas.height)
      
      for (let inx = 0; inx < stars.length; inx++) {
        const star = stars[inx]
        context.fillStyle = star.color
        context.shadowColor = star.shadowColor
        context.shadowBlur = star.shadowBlur
        context.fillRect(star.x, star.y, star.size, star.size)
      }
      if (positionChangeFlag) {
        setInterval(() => {
          moveStars(stars)
        }, 50);
        positionChangeFlag = false
      }
      requestAnimationFrame(draw)
    }
    draw()
  }
}
</script>

<style lang="sass">
.space
  width: 100vw
  height: 100vh
  position: fixed
  left: 0
  top: 0
  z-index: -1
  #globalBackground
    width: 100%
    height: 100%
</style>
