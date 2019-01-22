
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
      about: "Hello! I am a full-stack web developer. Previously engaged in the development and administration of systems based on 1C. But I left this area to more eloquent programming languages like python, ruby and java script (in particular, I love the vuejs framework). I would like to learn and work with golang."      
    }
  },
}
</script>


<style lang="sass">
.about
  // flex-wrap: wrap
  display: flex
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

@media screen and (max-width: 1400px)
  .about
    .content
      .body
@media screen and (max-width: 1024px)
  .about
    .content
      .body
@media screen and (max-width: 900px)
  .about
    flex-direction: column !important
    width: 100%
    .content
      order: 1
      width: 100%
      padding: 0
      margin: 0
    .game
      order: 2
      padding: 0
      margin: 0
      width: 100%
      margin-top: 2rem
      // height: 100%
@media screen and (max-width: 425px)
  .about
    .game
      margin-left: -1rem
    .content
      .body
        width: 90%

</style>
