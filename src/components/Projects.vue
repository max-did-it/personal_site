

<script>
// import template from '../templates/projects.slim'
import imageProject1 from '../assets/project1.svg'
import imageProject2 from '../assets/project2.svg'
import imageProject3 from '../assets/project3.svg'

import project from './Project'
export default {
  components: {
    'project': project
  },
  // mixins: [template],
  data () {
    return {
      projects: [
      {
        image: imageProject3,
        header: 'паркинг-онлайн',
        href: 'https://паркинг-онлайн.рф',
        description: 'Owner, ProjectManager, Teamlead, Backend',
        tags: ["RoR", "Ruby"]
      },
      {
        image: imageProject2,
        header: 'dominantafashion',
        href: 'http://dominantafashion.ru',
        description: 'ProjectManager, Fullstack',
        tags: ["Vue", "RoR", "Ruby"]
      },
      {
        image: imageProject1,
        header: 'code note',
        href: '',
        description: 'Owner, Fullstack',
        tags: ["RoR", "Vue", "Grape"]
      }
    ]
    }
  },
}
</script>
<style lang="sass" scoped>
.projects
  // margin-left: 15%
  width: calc(100% - 15%)
  .content
    .body
      display: flex
      flex-wrap: wrap
      align-items: flex-start
      align-content: flex-start

</style>

<template>
<div class="projects" id="projects">
  <div class="content">
      <header>
        <h1 class="title">
          Projects
        </h1>
      </header>
      <div class="body">
        <project v-for="p in projects" :projectData="p"></project>
      </div>
  </div>
</div>
</template>
