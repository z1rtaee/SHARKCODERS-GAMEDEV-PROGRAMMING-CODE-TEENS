embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"chao.png\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/sprites/tiles/PLAYER.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1.0
  }
  rotation {
    z: 0.0026570067
    w: 0.9999965
  }
  scale {
    y: 0.975944
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"ground\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1.0\n"
  "      y: -11.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 15.771991\n"
  "  data: 4.922074\n"
  "  data: 6.7948904\n"
  "}\n"
  ""
}
