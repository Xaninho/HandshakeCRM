export function useNavigationMenu() {
    
    const navigationMenu = () => {
      return [
        {
          label: 'Home',
          items: [{
            label: 'Dashboard', to: '/dashboard', icon: 'pi pi-fw pi-home',
          }],
        },
        {
          label: 'Clients',
          items: [
            { label: 'Contacts', icon: 'pi pi-fw pi-id-card', to: '/dashboard/clients/dashboardContacts' },
            { label: 'Companies', icon: 'pi pi-fw pi-truck', to: '/dashboard/clients/dashboardCompanies' },
          ],
        },
        {
          label: 'Admin',
          items: [
            { label: 'Manage Players', icon: 'pi pi-fw pi-users', to: '/dashboard/admin/dashboardPlayers' },
          ],
        } 
      ]
    }
  
    return { navigationMenu }
  }