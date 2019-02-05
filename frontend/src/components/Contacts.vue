
<script>
import template from '../templates/contacts.slim'
import icoGit from '../assets/Contacts/gitlab.svg'
import icoVk from '../assets/Contacts/vk.svg'
import icoTel from '../assets/Contacts/telegram.svg'
import icoMail from '../assets/Contacts/mail.svg'
import jq from 'jquery'
export default {
  mixins: [template],
  data () {
    return {
      popFormFlagLoading: false,
      popFormFlagSuccess: false,
      popFormFlagFailed: false,
      form: {
        name: '',
        email: '',
        message: '',

      },
      links: [
        {
          ico: icoGit,
          text: 'SiberianPanda',
          href: 'https://gitlab.com/SiberianPanda',
          target: '_blank'
        },
        {
          ico: icoVk,
          text: 'max_bez_l',
          href: 'https://vk.com/max_bez_l',
          target: '_blank'
        },
        {
          ico: icoTel,
          text: '@SiberiaThisAfterUral',
          href: '#contacts',
          target: ''
        },
        {
          ico: icoMail,
          text: 'max.bez.l@mail.ru',
          href: 'mailto:max.bez.l@mail.ru',
          target: ''
        },                        
      ]
    }
  },
  methods: {
    send_form(e) {
      let that = this
      that.popFormFlagLoading = true
      jq.ajax({
        type: "POST",
        url: "/callback",
        data: JSON.stringify({
          name: that.form.name,
          email: that.form.email,
          message: that.form.message
        }),
        contentType: "application/json;charset=utf-8",
        dataType: "json",
        success: function (response) {
          
          // console.log("Fuck yeah");
          // console.log(response.success);
          that.popFormFlagLoading = false
          if (response.success) {
            that.form.name = ''
            that.form.email = ''
            that.form.message = ''
            that.popFormFlagSuccess = true
          } else {
            that.popFormFlagFailed = true
          }
        },
        error: function (response) {
          that.popFormFlagLoading = false
          that.popFormFlagFailed = true
        },
        
      });

      return false
    }
  },
}
</script>

<style lang="sass">
.lds-ring 
  display: inline-block 
  position: relative 
  width: 64px 
  height: 64px 
  div 
    box-sizing: border-box 
    display: block 
    position: absolute 
    width: 51px 
    height: 51px 
    margin: 6px 
    border: 6px solid #fff
    border-radius: 50% 
    animation: lds-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite 
    border-color: #fff transparent transparent transparent 
    &:nth-child(1) 
      animation-delay: -0.45s 
    &:nth-child(2) 
      animation-delay: -0.3s 
    &:nth-child(3) 
      animation-delay: -0.15s 
@keyframes lds-ring 
  0%
    transform: rotate(0deg) 
  100% 
    transform: rotate(360deg) 
.contacts
  // background: white !important

  .form_submited
    position: fixed
    top: 0
    left: 0
    display: flex
    justify-content: center
    align-items: center
    z-index: 10000
    width: 100vw
    height: 100vh
    .content
      border: solid #C0A062
      padding: 1.5rem 
      display: flex
      justify-content: center
      align-items: center
      flex-direction: column
      background: black
      margin: 1rem
      .header
        padding-bottom: 1rem
      .text
        font-size: 1.5rem
      .close
        margin: 2rem
        padding: 0 1.5rem
        color: black
        border-radius: 1.5rem
        text-align: center
        font-weight: bold
        font-size: 1.5rem
        background: #C0A062
        cursor: pointer
  .content
    .body
      .links
        display: flex
        // justify-content: flex-start
        align-items: center
        flex-direction: column
        flex-wrap: wrap
        width: 50%
        height: 100%
        .link
          display: flex
          width: 50%
          height: 3rem
          justify-content: flex-start
          align-items: center
          flex-wrap: nowrap
          margin-bottom: 2rem
          .icon
            width: 3rem
            padding-right: 2rem
          .descript
            text-decoration: none
            border: none
            outline: none
            color: #c0a062
            font-size: 1.5rem
      .callback
        display: flex
        flex-direction: row
        flex-wrap: wrap
        padding: 0 5rem
        width: 50%
          // width: 70%
        display: flex
        justify-content: space-between
        input
          
        label
          box-sizing: padding-box
          padding-right: 5%
          width: 30%
          text-align: right
          justify-self: flex-end
          // align-self: flex-end
          color: #c0a062
          font-size: 1.3rem
          margin-left: auto
        .input
          margin-bottom: 2rem
          box-sizing: border-box
          color: white
          padding: 0.5rem
          height: 3rem
          font-size: 1.2rem
          width: 50%
          justify-self: flex-end
          align-self: end
          background: black
          outline: none
          border: 2px solid #c0a062
          &::placeholder
            color: rgba(192, 160, 98, 0.5)
        textarea
          height: 15rem !important
        input[type="submit"]
          margin: 0 auto
          width: 40%
          padding: 0.5rem 1.5rem
          background: #c0a062
          color: black
          border: none
          // font-weight: bold
          font-size: 1.05rem
          border-radius: 1.5rem
@media screen and (max-width: 1200px)
  .contacts
    .content
      .body
        .links
          align-items: flex-start
        .callback
          padding-right: 1rem
          padding-left: 2rem
          label
            display: none
            padding-right: 0
            margin-right: 1rem
          .input
            width: 75%
          input[type="submit"]
            width: 75%
            margin: 0
            
@media screen and (max-width: 900px)
  .contacts
    .content
      .body
        .callback
          .input
            width: 100% !important
            margin-right: 1rem
          input[type="submit"]
            width: 60% !important
@media screen and (max-width: 768px)
  .contacts
    .content
      .body
        flex-direction: column
        align-self: flex-start
        .links
          order: 1
          padding: 0
          margin: 0
          
          // width: calc(40% - 5rem)
          // padding-left: 5rem
        .callback
          order: 2
          flex-direction: column
          margin: 2rem 0
          padding: 0
          align-items: center
          width: 80%
          input , textarea
            box-sizing: border-box
            margin: 0 auto
            width: 60%
          input[type="submit"]
            width: 50%

</style>
