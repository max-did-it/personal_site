<template >
  <canvas id="spacegame"></canvas>
</template>

<script>
// import template from '../templates/game.slim'
import $ from 'vue-jquery'
import ShipFrame1 from "../assets/Game/spaceship-Frame1.png"
import ShipFrame2 from "../assets/Game/spaceship-Frame2.png"
import MeteorFrame1 from "../assets/Game/Meteor-Frame1.png"
import MeteorFrame2 from "../assets/Game/Meteor-Frame2.png"
import MeteorFrame3 from "../assets/Game/Meteor-Frame3.png"
import MeteorFrame4 from "../assets/Game/Meteor-Frame4.png"
import MeteorFrame5 from "../assets/Game/Meteor-Frame5.png"
import BulletFrame from "../assets/Game/Shot-Frame.png"
import PlayButtonFrame from "../assets/Game/Play-Button.png"
import GameFont from "../assets/Game/GameFont.ttf"
import pixel from "../assets/Game/pixel.png"
export default {
  mounted() {
    let canvas = document.getElementById("spacegame")
    canvas.width = window.innerWidth/3
    canvas.height = window.innerHeight
    let context = canvas.getContext("2d")
    let requestAnimationFrame = window.requestAnimationFrame || window.mozRequestAnimationFrame ||
      window.webkitRequestAnimationFrame || window.msRequestAnimationFrame
    let startMeteorGenerator = true
    let font = new FontFace('PixelFont', GameFont)
    console.log(font);
    
    font.onload = function (event) {
      document.fonts.add(font)
      context.font='15px PixelFont'
      console.log("here");
      
    }
    class Picture {
      constructor(game, image, ratio, x, y) {
        this.game = game
        this.canvas = game.canvas
        if (!image && true) {
          this.image = new Image()
          this.image.src = pixel
        } else {
          this.image = image
        }
        this.ratioSize = ratio || 0.18

        let object = this
        this.image.onload = function () {          
          object.initSize(object)
          object.initXY(object)
      } 
      }

      initSize(obj) {
          let object = this || obj
          object.width = object.canvas.width*object.ratioSize
          object.height = object.width*(object.image.height/object.image.width)
      }

      initXY(obj, x, y) {
        let object = this || obj 
        object.x = x || (object.game.canvas.width/2 - object.width/2)
        object.y = y || (object.game.canvas.height/2 - object.height/2)
      }


    }

    class Bullet extends Picture{ 
      constructor(game, parent, name ,speed) {
        super(game)
        this.name = name
        this.parent = parent
        this.speed = speed || 8
        this.ratioSize = parent.ratioSize/3
        this.image.src = BulletFrame
        this.path = setInterval(() => {
          this.y -= this.speed
          if ((this.y  + this.height) < 0) {this.stopBullet()}
        }, 20);
      }

      initXY(obj) {
        let object = this || obj
        object.x = object.parent.x + object.parent.width/3
        object.y = object.parent.y
      }
      
      stopBullet() {
        clearInterval(this.path)
        this.game.removeBullet(this.name)
      }
    }

    class Meteor extends Picture {
      constructor(game, name) {
        super(game)
        this.name = name
        this.currentFrameinx = 0
        frames = [
          MeteorFrame1,
          MeteorFrame2,
          MeteorFrame3,
          MeteorFrame4,
          MeteorFrame5
        ]
        this.framesCount = frames.length - 1
        this.images = []
        for (const inx in frames) {
          if (frames.hasOwnProperty(inx)) {
            const frame = frames[inx];
            let image = new Image
            image.src = frame
            this.images.push(image)
          }
        }
        this.setCurrentImage()
        this.flag = true
        this.path = setInterval(() => {
          this.y += 1

          // Check battleground bottom border
          if ((this.y - this.height) > this.canvas.height) {this.stopMeteor()}

          // Check collision with bullet
          for (const inx in this.game.bullets) {
            if (this.game.bullets.hasOwnProperty(inx)) {
              const bullet = this.game.bullets[inx];
              
              let bx1 = bullet.x, 
                  bx2 = bullet.x + bullet.width,
                  by1 = bullet.y,
                  by2 = bullet.y + bullet.height
              
              let mx1 = this.x,
                  mx2 = this.x + this.width,
                  my1 = this.y,
                  my2 = this.y + this.height
              
              let conditionX = ((bx1 >= mx1) && (bx1 <= mx2)) || ((bx2 <= mx2) && (bx2 >= mx1))
              let conditionY = ((by1 <= my2 ) && (by1 >= my1)) || ((by2 <= my2) && (by2 >= my1))
              if (conditionY && conditionX) {
                this.game.removeBullet(inx)
                if (this.currentFrameinx < this.framesCount) { 
                  this.currentFrameinx += 1
                  this.setCurrentImage()
                } else {
                  this.stopMeteor(inx)
                  this.game.score += 10
                }
              }
            }
          }
        }, 1);
      }
      setCurrentImage() {
        this.image = this.images[this.currentFrameinx]
        
      }

      initXY(obj) {
        let object = this || obj
        object.x = Math.floor(Math.random()*object.canvas.width) + 1
        object.y = 0
      }

      stopMeteor() {
        clearInterval(this.path)
        this.path = undefined
        this.game.removeMeteor(this.name)
      }
    }

    class Spaceship extends Picture {

      constructor(game) {
        super(game)
        this.currentFrame = 0
        this.image.src = ShipFrame1
        this.imageFrame2 = new Image
        this.imageFrame2.src = ShipFrame2
        this.ratioSize = 0.2

        let object = this

        this.canvas.addEventListener("mousemove",function (event) {
          object.setPosition(event.clientX - object.canvas.offsetLeft)
        }, false)

        this.canvas.addEventListener("click", function (event) {
          object.shoot()
        })

        setInterval(() => {
          this.image = [this.imageFrame2, this.imageFrame2 = this.image][0]
        }, 100);
      }

      initXY(obj) {
        let object = this || obj
        object.x = object.canvas.width/2 - object.width/3
        object.y = object.canvas.height - object.height - object.canvas.height*0.1
      }

      setPosition(x) {
        this.x  = x
      }

      shoot() {
        this.game.addBullet()
      }
    }


    class Game {
      constructor(canvas, context) {
        this.canvas = canvas
        this.context = context
        this.meteors = {}
        this.meteorsCount = 0
        this.bulletsCount = 0
        this.bullets = []
        this.playing = false
        this.score = 0
        let image = new Image()
        image.src = PlayButtonFrame
        this.playButton = new Picture(this, image, 0.3)
        let that = this
        this.canvas.addEventListener("click", function () {
          if (!that.playing) {
            that.play()
          }
        })
      }

      addMeteor() {
        this.meteorsCount += 1
        name = "Meteor" + this.meteorsCount
        this.meteors[name] = new Meteor(this, name)
      }

      addBullet() {
        this.bulletsCount += 1
        name = "Bullet" + this.bulletsCount
        this.bullets[name] = new Bullet(this, this.ship, name)
      }

      removeMeteor(name) {
        // this.meteorsCount -= 1
        this.meteors[name].x = undefined
        this.meteors[name].y = undefined
        delete this.meteors[name]
      }

      removeBullet(name) {
        this.bullets[name].x = undefined
        this.bullets[name].y = undefined
        delete this.bullets[name]
      }

      start() {
        this.ship = new Spaceship(this)
        draw()
        this.play()
      }

      play() {
        canvas.style.cursor = "none"
        this.playing = true

      }

      pause() {
        canvas.style.cursor = "auto"
        this.playing = false
        this.context.drawImage(
          this.playButton.image,
          this.playButton.x,
          this.playButton.y,
          this.playButton.height,
          this.playButton.width
        )

      }
    }


    function drawObjects(objects) {
      for (const inx in objects) {
        if (objects.hasOwnProperty(inx)) {
          const object = objects[inx]
          game.context.drawImage(object.image, object.x , object.y, object.width, object.height)
        }
      }
    }

    function draw() {
      if (game.playing) {
        if (startMeteorGenerator) {
          let blurFlag = true
          window.onblur = function () { game.pause() }
          setInterval(() => {
            if (blurFlag) {
              game.addMeteor();
            }
          }, 3000);
          startMeteorGenerator = false
        }
        game.context.fillStyle = "rgb(0,0,0)"
        game.context.fillRect (0, 0, game.canvas.width, game.canvas.height)
        drawObjects(game.meteors)
        drawObjects(game.bullets)
        game.context.fillStyle = "rgb(255,255,255)";
        game.context.drawImage(game.ship.image, game.ship.x , game.ship.y, game.ship.width, game.ship.height)
        game.context.fillText(game.score, 5, game.canvas.height - 5, game.canvas.width);
        
      }
        requestAnimationFrame(draw)
    }

    let game = new Game(canvas, context)
    game.start()
    document.game = game
  }
}
</script>
