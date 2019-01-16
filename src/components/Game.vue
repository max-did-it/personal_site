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
import BulletFrame from "../assets/Game/Shot-Frame.png"
import pixel from "../assets/Game/pixel.png"
export default {
  mounted() {
    let canvas = document.getElementById("spacegame")
    canvas.style.cursor = "none"
    canvas.width = window.innerWidth/3
    canvas.height = window.innerHeight
    let context = canvas.getContext("2d")


    class Picture {
      constructor(game) {
        this.game = game
        this.canvas = game.canvas
        this.image = new Image()
        this.image.src = pixel
        this.ratioSize = 0.1

        let object = this
        this.image.onload = function () {          
          object.initSize(object)
          object.initXY(object)
          draw()
      } 
      }

      initSize(obj) {
          let object = this || obj
          object.width = object.canvas.width*object.ratioSize
          object.height = object.width*(object.image.height/object.image.width)
      }

      initXY(obj) {
        let object = this || obj 
        object.x = 0
        object.y = 0
      }


    }

    class Bullet extends Picture{ 
      constructor(game, parent, inx ,speed) {
        super(game)
        this.inx = inx
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
        console.log("Interval " + this.path);
        
        clearInterval(this.path)
        this.game.removeBullet(this.inx)
      }
    }

    class Meteor extends Picture {
      constructor() {
        super(game)
        while (true) {
          
          let name = "Metheor" + Math.round(Math.random() * 40)
          console.log("name selecting: " + name);
          if (!game.meteors[name] ) break
        }
        this.name = name
        this.currentFrameinx = 0
        frames = [
          MeteorFrame1,
          MeteorFrame2,
          MeteorFrame3,
          MeteorFrame4
        ]
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

        this.path = setInterval(() => {
          this.y += 1
          if ((this.y + this.height) > this.canvas.height) {this.stopMeteor()}
          for (const inx in this.game.bullets) {
            if (this.game.bullets.hasOwnProperty(inx)) {
              const bullet = this.game.bullets[inx];
              if (bullet.x >= (this.x) && bullet.x <= (this.x + this.width)) {
                if (bullet.y >= (this.y) && bullet.y <= (this.y + this.height)) {
                  this.currentFrameinx += 1
                  this.setCurrentImage()
                }
              }
            }
          }
        }, 10);
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
        object.x = object.canvas.width/2 - object.width/2
        object.y = object.canvas.height - object.height - 10
      }

      setPosition(x) {
        this.x = x - this.width/2 
      }

      shoot() {
        this.game.addBullet()
      }
    }


    class Game {
      constructor(canvas, context) {
        this.canvas = canvas
        this.context = context
        this.ship = new Spaceship(this)
        this.meteors = {}
        this.bullets = []
      }

      addMeteor() {
        console.log("Meteor created");
        
        let a = new Meteor(this)
        this.meteors[a.name] = a
      }

      addBullet() {
        
        this.bullets.push(new Bullet(this, this.ship, this.bullets.length))
        console.log(this.bullets);
      }

      removeMeteor(name) {
        delete this.meteors[name]
      }

      removeBullet(inx) {
        this.bullets.shift()
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

    let flag = true
    function draw() {
        if (flag) {
          setInterval(() => {
            game.addMeteor();
          }, 3000);
          flag = false
        }
        game.context.fillStyle = "rgb(0,0,0)"
        game.context.fillRect (0, 0, game.canvas.width, game.canvas.height)
        drawObjects(game.meteors)
        drawObjects(game.bullets)
        game.context.drawImage(game.ship.image, game.ship.x , game.ship.y, game.ship.width, game.ship.height)
        requestAnimationFrame(draw)
    }

    let game = new Game(canvas, context)

  }
}
</script>
