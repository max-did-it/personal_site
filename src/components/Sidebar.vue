
<script>
import template from '../templates/sidebar.slim'
export default {
  mixins: [template],
  data () {
    return {
    }
  },
  mounted() {
    let history = window.history
    if ('scrollRestoration' in history) {
      history.scrollRestoration = 'manual';
    }
    window.addEventListener("scroll", function () {
      let blocks = document.getElementsByClassName("container")
      for (const inx in blocks) {
        if (blocks.hasOwnProperty(inx)) {
          const block = blocks[inx];
          let rect = block.getBoundingClientRect()
          if (Math.abs(rect.top) >= 0 && Math.abs(rect.top) < window.innerHeight*0.5) {
            
            if (window.history.state != block.id) {
              window.history.pushState(block.id, block.id, "/#"+block.id)
              document.getElementsByClassName("sidebar-elem-active")[0].classList.remove("sidebar-elem-active")
              document.getElementsByClassName(block.id + "-nav")[0].classList.add("sidebar-elem-active")
            }
            return true
          }
        }
      }
    }, false)
    window.scroll()
  }
}
</script>

<style lang="sass">
.sidebar
  margin: 7.5% 0 0 1rem
  border-left: 1px solid #c0a062
  position: fixed
  display: flex
  flex-direction: column
  justify-content: space-between
  height: 20rem
  padding: 1rem 0
  .sidebar-elem-active
    .circle
      background: black !important
      display: flex
      justify-content: center
      align-items: center
      font-size: 2rem
      font-weight: bold
      color: #c0a062
      letter-spacing: 0
      p
        position: relative
        top: -0.65rem
        padding: 0
        margin: 0
      p:after
        content: "."
  .sidebar-elem
    display: flex
    position: relative
    list-style-type: none
    justify-content: center
    align-items: center
    .circle
      position: absolute
      left: -0.65rem
      width: 1rem
      height: 1rem
      background: #c0a062
      border: 2px solid #c0a062
      border-radius: 1rem
      transition: all 0.5s 0s ease
    a
      line-height: 0.8rem
      width: 5rem
      text-align: center
      margin-left: 1rem
      padding: 3px 5px
      border-radius: 1.5rem
      border: 2px solid #c0a062
      text-decoration: none
      background: #c0a062
      color: black
      text-transform: lowercase
      transition: all 0.5s 0s ease
      &:hover
        background: none
        color: #c0a062
      
</style>
