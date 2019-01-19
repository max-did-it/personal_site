<script>
export default {
  props: ['projectData'],
  mounted() {
    console.log("Given data " + this.projectData);
  }
}
</script>

<template>
  <div class="project">
    <div class="project-header">
      <a :href="projectData.href" target="_blank">
        <h3>
          {{ projectData.header }}
        </h3>
      </a>
      <div class="arrow"></div>
    </div>
    <div class="project-img">
      <img :src="projectData.image" alt="" srcset="">
    </div>
    <div class="project-descript">
      <p>
        {{ projectData.description }}
      </p>
      <span>
        {{projectData.description}}
      </span>
    </div>
    <ul class="project-tags">
      <li class="tag" v-for="tag in projectData.tags">
        {{ tag }}
      </li>
    </ul>
  </div>
</template>

<style lang="sass">
  .project
    border: 1px solid  #C0A062
    width: 15vw
    min-width: 250px
    min-height: 20vw
    margin: 3% 0 0 auto
    margin-left: auto
    padding: 0rem 3rem 0 3rem
    box-sizing: content-box
    transition: height 1s linear
    .project-img
      img
        width: 100%
        box-sizing: border-box
      // padding: 3rem
    .project-descript
      span
        position: absolute
        display: none
        background: #000
        border: 1px solid #C0A062
        padding: 1rem 2rem 1rem 1rem
        max-width: 10vw 
      p
        // width: 15rem
        text-overflow: ellipsis
        white-space: nowrap
        overflow: hidden
        cursor: pointer
        &:hover ~ span
          display: block
          background: rgba(0,0,0,1)
          z-index: 1000
    .project-tags
      position: relative
      bottom: 0rem
      right: -2.5rem
      float: right
      display: flex
      .tag
        border: 1px solid #C0A062
        border-radius: 1rem
        padding: 0.5rem 1rem 0.5rem 1rem
        color: #C0A062
        list-style-type: none
        margin: 0.5rem
    .project-header
      position: relative
      top: 0rem
      left: -3rem
      background: #C0A062
      margin-bottom: 1rem
      transition: width 500ms linear
      width: 100%
      &:hover
        width: calc(100% + 6rem)
      .arrow
        left: 100%
        top: 50%
        border: solid transparent
        content: " "
        height: 0
        width: 0
        position: absolute
        pointer-events: none
        border-color: rgba(136, 183, 213, 0)
        border-left-color: #C0A062
        border-width: 1.1rem
        margin-top: -1.1rem
      a
        text-overflow: ellipsis
        white-space: nowrap
        overflow: hidden
        display: block
        text-decoration: none
        color: black
        padding: 0.25rem 1rem 0.25rem 1rem
        width: 100%
        h3
          padding: 0
          margin: 0
</style>
