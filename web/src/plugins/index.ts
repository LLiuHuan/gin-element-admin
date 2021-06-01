import {createApp} from 'vue'

/**
 * @description 加载所有 Plugins
 * @param  {ReturnType<createApp>} app 整个应用的实例
 */
export function loadAllPlugins(app: ReturnType<typeof createApp>) {
    const components = import.meta.globEager('./*.ts')
    Object.keys(components).forEach(path => {
        const fileName = path.replace(/(.*\/)*([^.]+).*/ig, "$2")
        // console.log(components[path].default);
        if (typeof components[path].default === 'function') {
            if (fileName !== './index.ts') components[path].default(app)
        }
    })
}
