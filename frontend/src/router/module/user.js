const userRoutes = [
    {
        path: '/register',
        name: 'register',
        component: () => import('@/views/register/userRegister.vue'),
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/login/userLogin.vue'),
    },
    {
        path: '/profile',
        name: 'profile',
        meta: {
            auth: true,
        },
        component: () => import('@/views/profile/layout_Profile.vue'),
    },
];

export default userRoutes;
