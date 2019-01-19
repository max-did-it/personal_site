<script>
import template from './templates/index.slim'
import SideBar from './components/Sidebar'
import Dog from './components/Dog'
import About from './components/About'
import Projects from './components/Projects'
import Contacts from './components/Contacts'
export default {
  mixins: [template],
  components: {
    'sidebar': SideBar,
    'dog': Dog,
    'about': About,
    'contacts': Contacts,
    'projects': Projects
  },
  name: 'App',
  mounted() {
    window.scrollTo( 0,document.getElementById(window.history.state).getBoundingClientRect().top)
  }
}
</script>

<style lang="sass">
@import url('https://fonts.googleapis.com/css?family=K2D:400,700')
html
  scroll-behavior: smooth
body 
  font-family: 'K2D', sans-serif
  background: black
  padding: 0
  margin: 0

* 
  -moz-user-select: none
  -ms-user-select: none
  -o-user-select: none
  -webkit-user-select: none
.test
  height: 50vh
#app 
  position: relative
  -webkit-font-smoothing: antialiased
  -moz-osx-font-smoothing: grayscale
  margin: 0 10% 0 10%
  padding-top: 0
  color: white
  .container
    min-height: 100vh
    width: 100%
    display: flex
    flex-direction: column
    header
      text-align: right
      width: 100%
      height: 20vh
      display: flex
      flex-direction: column
      align-content: flex-end
      .subtitle
        justify-self: flex-start
        color: white
      .title
        color: #C0A062
        margin-top: auto
        margin-left: auto
        // width: 30vw
        font-size: 3rem
        font-weight: normal
        
        justify-self: self-end
    .content
      .body
        display: flex
        margin-top: 13%
</style>
