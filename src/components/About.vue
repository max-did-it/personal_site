
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
  flex-direction: row !important
  .game
    // padding: 20rem
    margin-left: 15%
    margin-right: 5%
    // border: 
    border: 1px dashed white
    border-top: none
    border-bottom: none
    width: 30%
    height: 100%
  .content
    width: 70%
    article
      text-align: right
      .text
        float: right
        margin-top: 13%
</style>
