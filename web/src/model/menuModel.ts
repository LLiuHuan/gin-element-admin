export interface menuObject {
    CreatedAt: string;
    ID: number;
    UpdatedAt: string;
    authoritys?: string;
    children?: string;
    component: string;
    hidden: boolean;
    menuId: string
    meta: meta;
    name: string;
    parameters: string[];
    parentId: string;
    path: string;
    sort: number;
}

interface meta {
keepAlive: boolean;
defaultMenu: boolean;
title: string;
icon: string;
closeTab: boolean;
}