import { acceptHMRUpdate, defineStore } from 'pinia'

export function updateTheme(themeName: string, themeColor: string) {
  const newValue = `https://cdn.jsdelivr.net/npm/primevue@3.15.0/resources/themes/${themeName}-${themeColor}/theme.css`
  return newValue
}

export const useThemeStore = defineStore('theme', {

  state: () => ({
    themeName: 'saga',
    themeColor: 'green',
    link: 'https://cdn.jsdelivr.net/npm/primevue@3.15.0/resources/themes/saga-green/theme.css',
  }),

  getters: {
    theme: (state) => {
      return `${state.themeName}-${state.themeColor}`
    },
    isDarkMode: state => state.themeName !== 'saga',
  },

  actions: {
    setDark() {
      this.themeName = 'arya'
      this.link = updateTheme(this.themeName, this.themeColor)
    },
    setLight() {
      this.themeName = 'saga'
      this.link = updateTheme(this.themeName, this.themeColor)
    },
    setColor(colorName: string) {
      this.themeColor = colorName
      this.link = updateTheme(this.themeName, this.themeColor)
    },
  },
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useThemeStore, import.meta.hot))