import Vue from 'vue'
import Router from 'vue-router'
import Logo from '@/components/Logo'
import Sign from '@/components/Sign'
import About from '@/components/About'
import Workspace from '@/components/Workspace'
import WorkspaceSearch from '@/components/Search'
import WorkspaceCreate from '@/components/EditNote'
import WorkspaceShow from '@/components/ShowNote'
import WorkspaceStorage from '@/components/StorageNotes'
import SearchProperties from '@/components/SearchProperties'
import NoteProperties from '@/components/NoteProperties'
import StorageProperties from '@/components/StorageProperties'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Logo',
      components
    },
    {
      path: '/sign',
      name: 'Sign',
      component: Sign
    },
    {
      path: '/workspace',
      name: 'Main',
      component: Workspace,
      children: [
        {
          path: 'search',
          components: {
            default: WorkspaceSearch,
            properties: SearchProperties
          }
        },
        {
          path: 'create',
          components: {
            default: WorkspaceCreate,
            properties: NoteProperties
          }
        },
        {
          path: 'storage',
          components: {
            default: WorkspaceStorage,
            properties: StorageProperties
          }
        },
        {
          path: 'edit',
          component: WorkspaceShow
        }
      ]
    },
    {
      path: '/about',
      name: 'About',
      component: About
    }
  ],
  mode: 'history'
})
