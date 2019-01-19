
<script>
import template from '../templates/about.slim'
import Game from './Game'
export default {
  components: {
    'game': Game
  },
  mixins: [template],
  data () {
    return {
    }
  },
}
</script>


<style lang="sass">
.about
  // flex-wrap: wrap
  flex-direction: row !important
  // border-box: padding-box
  width: 100%
  .game
    order: 1
    margin-right: 5%
    border: 1px dashed white
    border-top: none
    border-bottom: none
  .content
    order: 2
    width: 70%

    .body
      .text
        text-align: right
</style>
