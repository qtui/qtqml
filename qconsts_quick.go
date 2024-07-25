package qtqml

type QQuickItem__Flag = int

const QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = 1
const QQuickItem__ItemAcceptsInputMethod QQuickItem__Flag = 2
const QQuickItem__ItemIsFocusScope QQuickItem__Flag = 4
const QQuickItem__ItemHasContents QQuickItem__Flag = 8
const QQuickItem__ItemAcceptsDrops QQuickItem__Flag = 16

type QQuickItem__ItemChange = int

const QQuickItem__ItemChildAddedChange QQuickItem__ItemChange = 0
const QQuickItem__ItemChildRemovedChange QQuickItem__ItemChange = 1
const QQuickItem__ItemSceneChange QQuickItem__ItemChange = 2
const QQuickItem__ItemVisibleHasChanged QQuickItem__ItemChange = 3
const QQuickItem__ItemParentHasChanged QQuickItem__ItemChange = 4
const QQuickItem__ItemOpacityHasChanged QQuickItem__ItemChange = 5
const QQuickItem__ItemActiveFocusHasChanged QQuickItem__ItemChange = 6
const QQuickItem__ItemRotationHasChanged QQuickItem__ItemChange = 7
const QQuickItem__ItemAntialiasingHasChanged QQuickItem__ItemChange = 8
const QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = 9
const QQuickItem__ItemEnabledHasChanged QQuickItem__ItemChange = 10

type QQuickItem__TransformOrigin = int

const QQuickItem__TopLeft QQuickItem__TransformOrigin = 0
const QQuickItem__Top QQuickItem__TransformOrigin = 1
const QQuickItem__TopRight QQuickItem__TransformOrigin = 2
const QQuickItem__Left QQuickItem__TransformOrigin = 3
const QQuickItem__Center QQuickItem__TransformOrigin = 4
const QQuickItem__Right QQuickItem__TransformOrigin = 5
const QQuickItem__BottomLeft QQuickItem__TransformOrigin = 6
const QQuickItem__Bottom QQuickItem__TransformOrigin = 7
const QQuickItem__BottomRight QQuickItem__TransformOrigin = 8

type QQuickPaintedItem__RenderTarget = int

const QQuickPaintedItem__Image QQuickPaintedItem__RenderTarget = 0
const QQuickPaintedItem__FramebufferObject QQuickPaintedItem__RenderTarget = 1
const QQuickPaintedItem__InvertedYFramebufferObject QQuickPaintedItem__RenderTarget = 2

type QQuickPaintedItem__PerformanceHint = int

const QQuickPaintedItem__FastFBOResizing QQuickPaintedItem__PerformanceHint = 1

type QQuickView__ResizeMode = int

const QQuickView__SizeViewToRootObject QQuickView__ResizeMode = 0
const QQuickView__SizeRootObjectToView QQuickView__ResizeMode = 1

type QQuickView__Status = int

const QQuickView__Null QQuickView__Status = 0
const QQuickView__Ready QQuickView__Status = 1
const QQuickView__Loading QQuickView__Status = 2
const QQuickView__Error QQuickView__Status = 3

type QQuickWindow__CreateTextureOption = int

const QQuickWindow__TextureHasAlphaChannel QQuickWindow__CreateTextureOption = 1
const QQuickWindow__TextureHasMipmaps QQuickWindow__CreateTextureOption = 2
const QQuickWindow__TextureOwnsGLTexture QQuickWindow__CreateTextureOption = 4
const QQuickWindow__TextureCanUseAtlas QQuickWindow__CreateTextureOption = 8
const QQuickWindow__TextureIsOpaque QQuickWindow__CreateTextureOption = 16

type QQuickWindow__RenderStage = int

const QQuickWindow__BeforeSynchronizingStage QQuickWindow__RenderStage = 0
const QQuickWindow__AfterSynchronizingStage QQuickWindow__RenderStage = 1
const QQuickWindow__BeforeRenderingStage QQuickWindow__RenderStage = 2
const QQuickWindow__AfterRenderingStage QQuickWindow__RenderStage = 3
const QQuickWindow__AfterSwapStage QQuickWindow__RenderStage = 4
const QQuickWindow__NoStage QQuickWindow__RenderStage = 5

type QQuickWindow__SceneGraphError = int

const QQuickWindow__ContextNotAvailable QQuickWindow__SceneGraphError = 1

type QQuickWindow__TextRenderType = int

const QQuickWindow__QtTextRendering QQuickWindow__TextRenderType = 0
const QQuickWindow__NativeTextRendering QQuickWindow__TextRenderType = 1

type QSGAbstractRenderer__ClearModeBit = int

const QSGAbstractRenderer__ClearColorBuffer QSGAbstractRenderer__ClearModeBit = 1
const QSGAbstractRenderer__ClearDepthBuffer QSGAbstractRenderer__ClearModeBit = 2
const QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = 4

type QSGEngine__CreateTextureOption = int

const QSGEngine__TextureHasAlphaChannel QSGEngine__CreateTextureOption = 1
const QSGEngine__TextureOwnsGLTexture QSGEngine__CreateTextureOption = 4
const QSGEngine__TextureCanUseAtlas QSGEngine__CreateTextureOption = 8
const QSGEngine__TextureIsOpaque QSGEngine__CreateTextureOption = 16

type QSGGeometry__AttributeType = int

const QSGGeometry__UnknownAttribute QSGGeometry__AttributeType = 0
const QSGGeometry__PositionAttribute QSGGeometry__AttributeType = 1
const QSGGeometry__ColorAttribute QSGGeometry__AttributeType = 2
const QSGGeometry__TexCoordAttribute QSGGeometry__AttributeType = 3
const QSGGeometry__TexCoord1Attribute QSGGeometry__AttributeType = 4
const QSGGeometry__TexCoord2Attribute QSGGeometry__AttributeType = 5

type QSGGeometry__DataPattern = int

const QSGGeometry__AlwaysUploadPattern QSGGeometry__DataPattern = 0
const QSGGeometry__StreamPattern QSGGeometry__DataPattern = 1
const QSGGeometry__DynamicPattern QSGGeometry__DataPattern = 2
const QSGGeometry__StaticPattern QSGGeometry__DataPattern = 3

type QSGGeometry__DrawingMode = int

const QSGGeometry__DrawPoints QSGGeometry__DrawingMode = 0
const QSGGeometry__DrawLines QSGGeometry__DrawingMode = 1
const QSGGeometry__DrawLineLoop QSGGeometry__DrawingMode = 2
const QSGGeometry__DrawLineStrip QSGGeometry__DrawingMode = 3
const QSGGeometry__DrawTriangles QSGGeometry__DrawingMode = 4
const QSGGeometry__DrawTriangleStrip QSGGeometry__DrawingMode = 5
const QSGGeometry__DrawTriangleFan QSGGeometry__DrawingMode = 6

type QSGGeometry__Type = int

const QSGGeometry__ByteType QSGGeometry__Type = 5120
const QSGGeometry__UnsignedByteType QSGGeometry__Type = 5121
const QSGGeometry__ShortType QSGGeometry__Type = 5122
const QSGGeometry__UnsignedShortType QSGGeometry__Type = 5123
const QSGGeometry__IntType QSGGeometry__Type = 5124
const QSGGeometry__UnsignedIntType QSGGeometry__Type = 5125
const QSGGeometry__FloatType QSGGeometry__Type = 5126

type QSGImageNode__TextureCoordinatesTransformFlag = int

const QSGImageNode__NoTransform QSGImageNode__TextureCoordinatesTransformFlag = 0
const QSGImageNode__MirrorHorizontally QSGImageNode__TextureCoordinatesTransformFlag = 1
const QSGImageNode__MirrorVertically QSGImageNode__TextureCoordinatesTransformFlag = 2

type QSGMaterial__Flag = int

const QSGMaterial__Blending QSGMaterial__Flag = 1
const QSGMaterial__RequiresDeterminant QSGMaterial__Flag = 2
const QSGMaterial__RequiresFullMatrixExceptTranslate QSGMaterial__Flag = 6
const QSGMaterial__RequiresFullMatrix QSGMaterial__Flag = 14
const QSGMaterial__CustomCompileStep QSGMaterial__Flag = 16

type QSGNode__NodeType = int

const QSGNode__BasicNodeType QSGNode__NodeType = 0
const QSGNode__GeometryNodeType QSGNode__NodeType = 1
const QSGNode__TransformNodeType QSGNode__NodeType = 2
const QSGNode__ClipNodeType QSGNode__NodeType = 3
const QSGNode__OpacityNodeType QSGNode__NodeType = 4
const QSGNode__RootNodeType QSGNode__NodeType = 5
const QSGNode__RenderNodeType QSGNode__NodeType = 6

type QSGNode__Flag = int

const QSGNode__OwnedByParent QSGNode__Flag = 1
const QSGNode__UsePreprocess QSGNode__Flag = 2
const QSGNode__OwnsGeometry QSGNode__Flag = 65536
const QSGNode__OwnsMaterial QSGNode__Flag = 131072
const QSGNode__OwnsOpaqueMaterial QSGNode__Flag = 262144
const QSGNode__IsVisitableNode QSGNode__Flag = 16777216

type QSGNode__DirtyStateBit = int

const QSGNode__DirtySubtreeBlocked QSGNode__DirtyStateBit = 128
const QSGNode__DirtyMatrix QSGNode__DirtyStateBit = 256
const QSGNode__DirtyNodeAdded QSGNode__DirtyStateBit = 1024
const QSGNode__DirtyNodeRemoved QSGNode__DirtyStateBit = 2048
const QSGNode__DirtyGeometry QSGNode__DirtyStateBit = 4096
const QSGNode__DirtyMaterial QSGNode__DirtyStateBit = 8192
const QSGNode__DirtyOpacity QSGNode__DirtyStateBit = 16384
const QSGNode__DirtyForceUpdate QSGNode__DirtyStateBit = 32768
const QSGNode__DirtyUsePreprocess QSGNode__DirtyStateBit = 2
const QSGNode__DirtyPropagationMask QSGNode__DirtyStateBit = 50432

type QSGRendererInterface__GraphicsApi = int

const QSGRendererInterface__Unknown QSGRendererInterface__GraphicsApi = 0
const QSGRendererInterface__Software QSGRendererInterface__GraphicsApi = 1
const QSGRendererInterface__OpenGL QSGRendererInterface__GraphicsApi = 2
const QSGRendererInterface__Direct3D12 QSGRendererInterface__GraphicsApi = 3
const QSGRendererInterface__OpenVG QSGRendererInterface__GraphicsApi = 4

type QSGRendererInterface__Resource = int

const QSGRendererInterface__DeviceResource QSGRendererInterface__Resource = 0
const QSGRendererInterface__CommandQueueResource QSGRendererInterface__Resource = 1
const QSGRendererInterface__CommandListResource QSGRendererInterface__Resource = 2
const QSGRendererInterface__PainterResource QSGRendererInterface__Resource = 3

type QSGRendererInterface__ShaderType = int

const QSGRendererInterface__UnknownShadingLanguage QSGRendererInterface__ShaderType = 0
const QSGRendererInterface__GLSL QSGRendererInterface__ShaderType = 1
const QSGRendererInterface__HLSL QSGRendererInterface__ShaderType = 2

type QSGRendererInterface__ShaderCompilationType = int

const QSGRendererInterface__RuntimeCompilation QSGRendererInterface__ShaderCompilationType = 1
const QSGRendererInterface__OfflineCompilation QSGRendererInterface__ShaderCompilationType = 2

type QSGRendererInterface__ShaderSourceType = int

const QSGRendererInterface__ShaderSourceString QSGRendererInterface__ShaderSourceType = 1
const QSGRendererInterface__ShaderSourceFile QSGRendererInterface__ShaderSourceType = 2
const QSGRendererInterface__ShaderByteCode QSGRendererInterface__ShaderSourceType = 4

type QSGRenderNode__StateFlag = int

const QSGRenderNode__DepthState QSGRenderNode__StateFlag = 1
const QSGRenderNode__StencilState QSGRenderNode__StateFlag = 2
const QSGRenderNode__ScissorState QSGRenderNode__StateFlag = 4
const QSGRenderNode__ColorState QSGRenderNode__StateFlag = 8
const QSGRenderNode__BlendState QSGRenderNode__StateFlag = 16
const QSGRenderNode__CullState QSGRenderNode__StateFlag = 32
const QSGRenderNode__ViewportState QSGRenderNode__StateFlag = 64
const QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = 128

type QSGRenderNode__RenderingFlag = int

const QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = 1
const QSGRenderNode__DepthAwareRendering QSGRenderNode__RenderingFlag = 2
const QSGRenderNode__OpaqueRendering QSGRenderNode__RenderingFlag = 4

type QSGSimpleTextureNode__TextureCoordinatesTransformFlag = int

const QSGSimpleTextureNode__NoTransform QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 0
const QSGSimpleTextureNode__MirrorHorizontally QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 1
const QSGSimpleTextureNode__MirrorVertically QSGSimpleTextureNode__TextureCoordinatesTransformFlag = 2

type QSGTexture__WrapMode = int

const QSGTexture__Repeat QSGTexture__WrapMode = 0
const QSGTexture__ClampToEdge QSGTexture__WrapMode = 1
const QSGTexture__MirroredRepeat QSGTexture__WrapMode = 2

type QSGTexture__Filtering = int

const QSGTexture__None QSGTexture__Filtering = 0
const QSGTexture__Nearest QSGTexture__Filtering = 1
const QSGTexture__Linear QSGTexture__Filtering = 2

type QSGTexture__AnisotropyLevel = int

const QSGTexture__AnisotropyNone QSGTexture__AnisotropyLevel = 0
const QSGTexture__Anisotropy2x QSGTexture__AnisotropyLevel = 1
const QSGTexture__Anisotropy4x QSGTexture__AnisotropyLevel = 2
const QSGTexture__Anisotropy8x QSGTexture__AnisotropyLevel = 3
const QSGTexture__Anisotropy16x QSGTexture__AnisotropyLevel = 4
