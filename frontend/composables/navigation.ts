export function useNavigationMenu() {
    
    const navigationMenu = () => {
      return [
        {
          label: 'Home',
          items: [{
            label: 'Dashboard', to: '/', icon: 'pi pi-fw pi-home',
          }],
        },
        {
          label: 'PrimeVue',
          items: [
            { label: 'DataTable', icon: 'pi pi-fw pi-table', to: '/prime/datatable' },
            { label: 'Messages', icon: 'pi pi-fw pi-user-edit', to: '/prime/messages' },
            { label: 'Validation', icon: 'pi pi-fw pi-user-edit', to: '/prime/validation' },
          ],
        },
        {
          label: 'UI',
          items: [
            { label: 'UnoCSS', icon: 'pi pi-fw pi-user-edit', to: '/ui/uno' },
            { label: 'Icons', icon: 'pi pi-fw pi-user-edit', to: '/ui/icons' },
          ],
        } 
      ]
    }
  
    return { navigationMenu }
  }