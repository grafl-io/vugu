package vugu

// GENERATED FILE, DO NOT EDIT!  See renderer-js-script-maker.go

const jsHelperScript = "\n(function() {\n\n\tif (window.vuguRender) { return; } // only once\n\n    const opcodeEnd = 0         // no more instructions in this buffer\n    // const opcodeClearRefmap = 1 // clear the reference map, all following instructions must not reference prior IDs\n    const opcodeClearEl = 1 // clear the currently selected element\n    // const opcodeSetHTMLRef = 2  // assign ref for html tag\n    // const opcodeSetHeadRef = 3  // assign ref for head tag\n    // const opcodeSetBodyRef = 4  // assign ref for body tag\n    // const opcodeSelectRef = 5   // select element by ref\n\tconst opcodeRemoveOtherAttrs = 5 // remove any elements for the current element that we didn't just set\n    const opcodeSetAttrStr = 6  // assign attribute string to the current selected element\n    const opcodeSelectMountPoint = 7 // selects the mount point element and pushes to the stack - the first time by selector but every subsequent time it will reuse the element from before (because the selector may not match after it's been synced over, it's id etc), also make sure it's of this element name and recreate if so\n\t// const opcodePicardFirstChildElement = 8  // ensure an element first child and push onto element stack\n\t// const opcodePicardFirstChildText    = 9  // ensure a text first child and push onto element stack\n\t// const opcodePicardFirstChildComment = 10 // ensure a comment first child and push onto element stack\n\t// const opcodeSelectParent                   = 11 // pop from the element stack\n\t// const opcodePicardFirstChild = 12  // ensure an element first child and push onto element stack\n\n    const opcodeMoveToFirstChild     = 20 // move node selection to first child (doesn't have to exist)\n\tconst opcodeSetElement           = 21 // assign current selected node as an element of the specified type\n\t// const opcodeSetElementAttr       = 22 // set attribute on current element\n\tconst opcodeSetText              = 23 // assign current selected node as text with specified content\n\tconst opcodeSetComment           = 24 // assign current selected node as comment with specified content\n\tconst opcodeMoveToParent         = 25 // move node selection to parent\n\tconst opcodeMoveToNextSibling    = 26 // move node selection to next sibling (doesn't have to exist)\n\tconst opcodeRemoveOtherEventListeners  = 27 // remove all event listeners from currently selected element that were not just set\n\tconst opcodeSetEventListener     = 28 // assign event listener to currently selected element\n    const opcodeSetInnerHTML         = 29 // set the innerHTML for an element\n\n    // Decoder provides our binary decoding.\n    // Using a class because that's what all the cool JS kids are doing these days.\n    class Decoder {\n\n        constructor(dataView, offset) {\n            this.dataView = dataView;\n            this.offset = offset || 0;\n            return this;\n        }\n\n        // readUint8 reads a single byte, 0-255\n        readUint8() {\n            var ret = this.dataView.getUint8(this.offset);\n            this.offset++;\n            return ret;\n        }\n\n        // readRefToString reads a 64-bit unsigned int ref but returns it as a hex string\n        readRefToString() {\n            // read in two 32-bit parts, BigInt is not yet well supported\n            var ret = this.dataView.getUint32(this.offset).toString(16).padStart(8, \"0\") +\n                this.dataView.getUint32(this.offset + 4).toString(16).padStart(8, \"0\");\n            this.offset += 8;\n            return ret;\n        }\n\n        // readString is 4 bytes length followed by utf chars\n        readString() {\n            var len = this.dataView.getUint32(this.offset);\n            var ret = utf8decoder.decode(new DataView(this.dataView.buffer, this.dataView.byteOffset + this.offset + 4, len));\n            this.offset += len + 4;\n            return ret;\n        }\n\n    }\n\n    let utf8decoder = new TextDecoder();\n\n\twindow.vuguSetEventHandlerAndBuffer = function(eventHandlerFunc, eventBuffer) { \n\t\tlet state = window.vuguRenderState || {};\n        window.vuguRenderState = state;\n        state.eventBuffer = eventBuffer;\n        state.eventHandlerFunc = eventHandlerFunc;\n    }\n\n\twindow.vuguRender = function(buffer) { \n        \n        // NOTE: vuguRender must not automatically reset anything between calls.\n        // Since a series of instructions might get cut off due to buffer end, we\n        // need to be able to just pick right up with the next call where we left off.\n        // The caller decides when to reset things by sending the appropriate\n        // instruction(s).\n\n\t\tlet state = window.vuguRenderState || {};\n\t\twindow.vuguRenderState = state;\n\n\t\tconsole.log(\"vuguRender called\", buffer);\n\n\t\tlet bufferView = new DataView(buffer.buffer, buffer.byteOffset, buffer.byteLength);\n\n        var decoder = new Decoder(bufferView, 0);\n        \n        // state.refMap = state.refMap || {};\n        // state.curRef = state.curRef || \"\"; // current reference number (as a hex string)\n        // state.curRefEl = state.curRefEl || null; // current reference element\n        // state.elStack = state.elStack || []; // stack of elements as we traverse the DOM tree\n\n        // mount point element\n        state.mountPointEl = state.mountPointEl || null; \n\n        // currently selected element\n        state.el = state.el || null;\n\n        // specifies a \"next\" move for the current element, if used it must be followed by\n        // one of opcodeSetElement, opcodeSetText, opcodeSetComment, which will create/replace/use existing\n        // the element and put it in \"el\".  The point is this allow us to select nodes that may\n        // not exist yet, knowing that the next call will specify what that node is.  It's more complex here\n        // but makes it easier to generate instructions while walking a DOM tree.\n        // Value is one of \"first_child\", \"next_sibling\"\n        // (Parents always exist and so doesn't use this mechanism.)\n        state.nextElMove = state.nextElMove || null;\n\n        // keeps track of attributes that are being set on the current element, so we can remove any extras\n        state.elAttrNames = state.elAttrNames || {};\n\n        // map of positionID -> array of listener spec and handler function, for all elements\n        state.eventHandlerMap = state.eventHandlerMap || {};\n    \n        // keeps track of event listeners that are being set on the current element, so we can remvoe any extras\n        state.elEventKeys = state.elEventKeys || {};\n\n        instructionLoop: while (true) {\n\n            let opcode = decoder.readUint8();\n            \n            // console.log(\"processing opcode\", opcode);\n            // console.log(\"test_span_id: \", document.querySelector(\"#test_span_id\"));\n\n            switch (opcode) {\n\n                case opcodeEnd: {\n                    break instructionLoop;\n                }\n    \n                // case opcodeClearRefmap:\n                //     state.refMap = {};\n                //     state.curRef = \"\";\n                //     state.curRefEl = null;\n                //     break;\n\n                case opcodeClearEl: {\n                    state.el = null;\n                    state.nextElMove = null;\n                    break;\n                }\n        \n                // case opcodeSetHTMLRef:\n                //     var refstr = decoder.readRefToString();\n                //     state.refMap[refstr] = document.querySelector(\"html\");\n                //     break;\n\n                // case opcodeSelectRef:\n                //     var refstr = decoder.readRefToString();\n                //     state.curRef = refstr;\n                //     state.curRefEl = state.refMap[refstr];\n                //     if (!state.curRefEl) {\n                //         throw \"opcodeSelectRef: refstr does not exist - \" + refstr;\n                //     }\n                //     break;\n\n                case opcodeSetAttrStr: {\n                    let el = state.el;\n                    if (!el) {\n                        return \"opcodeSetAttrStr: no current reference\";\n                    }\n                    let attrName = decoder.readString();\n                    let attrValue = decoder.readString();\n                    el.setAttribute(attrName, attrValue);\n                    state.elAttrNames[attrName] = true;\n                    // console.log(\"setting attr\", attrName, attrValue, el)\n                    break;\n                }\n\n                case opcodeSelectMountPoint: {\n                    \n                    state.elAttrNames = {}; // reset attribute list\n                    state.elEventKeys = {};\n\n                    // select mount point using selector or if it was done earlier re-use the one from before\n                    let selector = decoder.readString();\n                    let nodeName = decoder.readString();\n                    // console.log(\"GOT HERE selector,nodeName = \", selector, nodeName);\n                    // console.log(\"state.mountPointEl\", state.mountPointEl);\n                    if (state.mountPointEl) {\n                        console.log(\"opcodeSelectMountPoint: state.mountPointEl already exists, using it\", state.mountPointEl, \"parent is\", state.mountPointEl.parentNode);\n                        state.el = state.mountPointEl;\n                        // state.elStack.push(state.mountPointEl);\n                    } else {\n                        console.log(\"opcodeSelectMountPoint: state.mountPointEl does not exist, using selector to find it\", selector);\n                        let el = document.querySelector(selector);\n                        if (!el) {\n                            throw \"mount point selector not found: \" + selector;\n                        }\n                        state.mountPointEl = el;\n                        // state.elStack.push(el);\n                        state.el = el;\n                    }\n\n                    let el = state.el;\n\n                    // make sure it's the right element name and replace if not\n                    if (el.nodeName.toUpperCase() != nodeName.toUpperCase()) {\n\n                        let newEl = document.createElement(nodeName);\n                        el.parentNode.replaceChild(newEl, el);\n\n                        state.mountPointEl = newEl;\n                        el = newEl;\n\n                    }\n\n                    state.el = el;\n\n                    state.nextElMove = null;\n\n                    break;\n                }\n\n                // case opcodePicardFirstChild: {\n\n            \t// \tlet nodeType = decoder.readUint8();\n                //     let data = decoder.readString();\n\n                //     let oldFirstChildEl = state.el.firstChild;\n\n                //     let newFirstChildEl = null;\n\n                //     let needsCreate = true;\n                //     if (oldFirstChildEl) {\n                //         // node types from Go are https://godoc.org/golang.org/x/net/html#NodeType\n                //         // whereas node types in DOM are https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType\n\n                //         // text\n                //         if (nodeType == 1 && oldFirstChildEl.nodeType == 3) {\n                //             needsCreate = false;\n                //         } else \n                //         // element\n                //         if (nodeType == 3 && oldFirstChildEl.nodeType == 1) {\n                //             needsCreate = false;\n                //         } else \n                //         // comment\n                //         if (nodeType == 4 && oldFirstChildEl.nodeType == 8) {\n                //             needsCreate = false;\n                //         }\n\n                //     }\n\n                //     if (needsCreate) {\n\n                //         switch (nodeType) {\n                //             case 1: {\n                //                 newFirstChildEl = document.createTextNode(data);\n                //                 break;\n                //             }\n                //             case 3: {\n                //                 newFirstChildEl = document.createElement(data);\n                //                 break;\n                //             }\n                //             case 4: {\n                //                 newFirstChildEl = document.createComment(data);\n                //                 break;\n                //             }\n                //         }\n    \n                //     }\n\n                //     if (newFirstChildEl) {\n                //         if (oldFirstChildEl) {\n                //             state.el.replaceChild(newFirstChildEl, oldFirstChildEl);\n                //         } else {\n                //             state.el.appendChild(newFirstChildEl);\n                //         }\n                //         state.el = newFirstChildEl;\n                //     } else {\n                //         state.el = oldFirstChildEl;\n                //     }\n\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildElement: {\n                //     // ensure an element first child and select\n\n                //     let el = state.el;\n                //     let nextEl = el.firstChild;\n                //     if (!nextEl) {\n                //         nextEl = \n                //     }\n                //     state.el = el;\n\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildText: {\n                //     // ensure a text first child and select\n                //     break;\n                // }\n\n                // case opcodePicardFirstChildComment: {\n                //     // ensure a comment first child and select\n                //     break;\n                // }\n\n                // remove any elements for the current element that we didn't just set\n                case opcodeRemoveOtherAttrs: {\n\n                    if (!state.el) {\n                        throw \"no element selected\";\n                    }\n\n                    if (state.nextElMove) {\n                        throw \"cannot call opcodeRemoveOtherAttrs when nextElMove is set\";\n                    }\n\n                    // build a list of attribute names to remove\n                    let rmAttrNames = [];\n                    for (let i = 0; i < state.el.attributes.length; i++) {\n                        if (!state.elAttrNames[state.el.attributes[i].name]) {\n                            rmAttrNames.push(state.el.attributes[i].name);\n                        }\n                    }\n\n                    // remove them\n                    for (let i = 0; i < rmAttrNames.length; i++) {\n                        state.el.attributes.removeNamedItem(rmAttrNames[i]);\n                    }\n\n                    break;\n                }\n\n                // move node selection to parent\n                case opcodeMoveToParent: {\n\n                    // if first_child is next move then we just unset this\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                    } else {\n                        // otherwise we actually move and also reset nextElMove\n                        state.el = state.el.parentNode;\n                        state.nextElMove = null;\n                    }\n\n                    break;\n                }\n\n                // move node selection to first child (doesn't have to exist)\n                case opcodeMoveToFirstChild: {\n\n                    // if a next move already set, then we need to execute it before we can do this\n                    if (state.nextElMove) {\n                        if (state.nextElMove == \"first_child\") {\n                            state.el = state.el.firstChild;\n                            if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.el = state.el.nextSibling;\n                            if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                        }\n                        state.nextElMove = null;\n                    }\n\n                    if (!state.el) { throw \"must have current selection to use opcodeMoveToFirstChild\"; }\n                    state.nextElMove = \"first_child\";\n\n                    break;\n                }\n                \n                // move node selection to next sibling (doesn't have to exist)\n                case opcodeMoveToNextSibling: {\n\n                    // if a next move already set, then we need to execute it before we can do this\n                    if (state.nextElMove) {\n                        if (state.nextElMove == \"first_child\") {\n                            state.el = state.el.firstChild;\n                            if (!state.el) { throw \"unable to find state.el.firstChild\"; }\n                        } else if (state.nextElMove == \"next_sibling\") {\n                            state.el = state.el.nextSibling;\n                            if (!state.el) { throw \"unable to find state.el.nextSibling\"; }\n                        }\n                        state.nextElMove = null;\n                    }\n\n                    if (!state.el) { throw \"must have current selection to use opcodeMoveToNextSibling\"; }\n                    state.nextElMove = \"next_sibling\";\n\n                    break;\n                }\n                \n                // assign current selected node as an element of the specified type\n                case opcodeSetElement: {\n                    \n                    let nodeName = decoder.readString();\n\n                    state.elAttrNames = {};\n                    state.elEventKeys = {};\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        if (newEl) { \n                            state.el = newEl; \n                            break; \n                        } else {\n                            newEl = document.createElement(nodeName);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        if (newEl) { \n                            state.el = newEl; \n                            break; \n                        } else {\n                            newEl = document.createElement(nodeName);\n                            // console.log(\"HERE1\", state.el);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact an element of the right type\n                    // and replace if not\n\n                    if (state.el.nodeType != 1 || state.el.nodeName.toUpperCase() != nodeName.toUpperCase()) {\n\n                        let newEl = document.createElement(nodeName);\n                        // throw \"stopping here\";\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n\n                    }\n\n                    break;\n                }\n\n                // assign current selected node as text with specified content\n                case opcodeSetText: {\n\n                    let content = decoder.readString();\n\n                    // console.log(\"in opcodeSetText 1\");\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        // console.log(\"in opcodeSetText 2\");\n                        if (newEl) { \n                            state.el = newEl; \n                            break;\n                        } else {\n                            let newEl = document.createTextNode(content);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            // console.log(\"in opcodeSetText 3\");\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        // console.log(\"in opcodeSetText 4\");\n                        if (newEl) { \n                            state.el = newEl; \n                            break; \n                        } else {\n                            let newEl = document.createTextNode(content);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            // console.log(\"in opcodeSetText 5\");\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact a node of the right type\n                    // and with right content and replace if not\n                    // console.log(\"in opcodeSetText 6\");\n\n                    if (state.el.nodeType != 3) {\n\n                        let newEl = document.createTextNode(content);\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n                        // console.log(\"in opcodeSetText 7\");\n\n                    } else {\n                        // console.log(\"in opcodeSetText 8\");\n                        state.el.textContent = content;\n                    }\n                    // console.log(\"in opcodeSetText 9\");\n\n                    break;\n                }\n\n                // assign current selected node as comment with specified content\n                case opcodeSetComment: {\n                    \n                    let content = decoder.readString();\n\n                    // handle nextElMove cases\n\n                    if (state.nextElMove == \"first_child\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.firstChild;\n                        if (newEl) { \n                            state.el = newEl; \n                            break; \n                        } else {\n                            let newEl = document.createComment(content);\n                            state.el.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove == \"next_sibling\") {\n                        state.nextElMove = null;\n                        let newEl = state.el.nextSibling;\n                        if (newEl) { \n                            state.el = newEl; \n                            break; \n                        } else {\n                            let newEl = document.createComment(content);\n                            // state.el.insertAdjacentElement(newEl, 'afterend');\n                            state.el.parentNode.appendChild(newEl);\n                            state.el = newEl;\n                            break; // we're done here, since we just created the right element\n                        }\n                    } else if (state.nextElMove) {\n                        throw \"bad state.nextElMove value: \" + state.nextElMove;\n                    }\n\n                    // if we get here we need to verify that state.el is in fact a node of the right type\n                    // and with right content and replace if not\n\n                    if (state.el.nodeType != 8) {\n\n                        let newEl = document.createComment(content);\n                        state.el.parentNode.replaceChild(newEl, state.el);\n                        state.el = newEl;\n\n                    } else {\n                        state.el.textContent = content;\n                    }\n\n                    break;\n                }\n\n                case opcodeSetInnerHTML: {\n\n                    let html = decoder.readString();\n\n                    if (!state.el) { throw \"opcodeSetInnerHTML must have currently selected element\"; }\n                    if (state.nextElMove) { throw \"opcodeSetInnerHTML nextElMove must not be set\"; }\n                    if (state.el.nodeType != 1) { throw \"opcodeSetInnerHTML currently selected element expected nodeType 1 but has: \" + state.el.nodeType; }\n\n                    state.el.innerHTML = html;\n\n                    break;\n                }\n\n                // remove all event listeners from currently selected element that were not just set\n                case opcodeRemoveOtherEventListeners: {\n                    this.console.log(\"opcodeRemoveOtherEventListeners\");\n\n                    let positionID = decoder.readString();\n\n                    // look at all registered events for this positionID\n                    let emap = state.eventHandlerMap[positionID] || {};\n                    // for any that we didn't just set, remove them\n                    let toBeRemoved = [];\n                    for (let k in emap) {\n                        if (!state.elEventKeys[k]) {\n                            toBeRemoved.push(k);\n                        }\n                    }\n\n                    // for each one that was missing, we remove from emap and call removeEventListener\n                    for (let i = 0; i < toBeRemoved.length; i++) {\n                        let f = emap[k];\n                        let k = toBeRemoved[i];\n                        let kparts = k.split(\"|\");\n                        state.el.removeEventListener(kparts[0], f, {capture:!!kparts[1], passive:!!kparts[2]});\n                        delete emap[k];\n                    }\n\n                    // if emap is empty now, remove the entry from eventHandlerMap altogether\n                    if (Object.keys(emap).length == 0) {\n                        delete state.eventHandlerMap[positionID];\n                    } else {\n                        state.eventHandlerMap[positionID] = emap;\n                    }\n\n                    break;\n                }\n            \n                // assign event listener to currently selected element\n                case opcodeSetEventListener: {\n                    let positionID = decoder.readString();\n                    let eventType = decoder.readString();\n                    let capture = decoder.readUint8();\n                    let passive = decoder.readUint8();\n\n                    if (!state.el) {\n                        throw \"must have state.el set in order to call opcodeSetEventListener\";\n                    }\n\n                    var eventKey = eventType + \"|\" + (capture?\"1\":\"0\") + \"|\" + (passive?\"1\":\"0\");\n                    state.elEventKeys[eventKey] = true;\n\n                    // map of positionID -> map of listener spec and handler function, for all elements\n                    //state.eventHandlerMap\n                    let emap = state.eventHandlerMap[positionID] || {};\n\n                    // register function if not done already\n                    let f = emap[eventKey];\n                    if (!f) {\n                        f = function(event) {\n                            // TODO: serialize event into the event buffer, somehow,\n                            // and keep track of the target element, also consider grabbing\n                            // the value or relevant properties as appropriate for form things\n                            state.eventHandlerFunc.call(null); // call with null this avoid unnecessary js.Value reference\n                        };    \n                        emap[eventKey] = f;\n\n                        this.console.log(\"addEventListener\", eventType);\n                        state.el.addEventListener(eventType, f, {capture:capture, passive:passive});\n                    }\n\n                    state.eventHandlerMap[positionID] = emap;\n\n                    this.console.log(\"opcodeSetEventListener\", positionID, eventType, capture, passive);\n                    break;\n                }\n            \n                // case opcodeSelectParent: {\n                //     // select parent\n                //     state.el = state.el.parentNode;\n                //     break;\n                // }\n\n                default: {\n                    console.error(\"found invalid opcode\", opcode);\n                    return;\n                }\n            }\n\n\t\t}\n\n\t}\n\n})()\n"