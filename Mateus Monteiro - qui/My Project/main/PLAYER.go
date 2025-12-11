components {
  id: "player"
  component: "/main/assets/scripts/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player.png\"\n"
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
    x: 0.28179303
    w: 0.9594752
  }
  scale {
    x: 0.938362
    y: 1.121056
    z: 0.707354
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"ground\"\n"
  "mask: \"inimigo\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 10.933333\n"
  "  data: 9.505306\n"
  "  data: 10.933333\n"
  "}\n"
  ""
}
