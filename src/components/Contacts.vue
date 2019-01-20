
<script>
import template from '../templates/contacts.slim'
import icoGit from '../assets/Contacts/gitlab.svg'
import icoVk from '../assets/Contacts/vk.svg'
import icoTel from '../assets/Contacts/telegram.svg'
import icoMail from '../assets/Contacts/mail.svg'
export default {
  mixins: [template],
  data () {
    return {
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
  }
}
</script>

<style lang="sass">
.contacts
  // background: white !important
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
        align-items: center
        .links
          order: 1
          width: 100%
          padding-left: 5rem
        .callback
          order: 2
          flex-direction: column
          margin: 2rem 0
          align-items: center
          width: 100%
          input , textarea
            margin: 0 auto
          input[type="submit"]
            width: 50%

</style>
