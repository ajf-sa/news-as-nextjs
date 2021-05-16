import HeaderContext from '../contexts/HeaderContext'
import { useState } from 'react'

function ContextWrapper({children, tags}) {
    const [menuItems] = useState(tags)

    return (
        <HeaderContext.Provider value={{menuItems}}>
            {children}
        </HeaderContext.Provider>
    )
}

export default ContextWrapper