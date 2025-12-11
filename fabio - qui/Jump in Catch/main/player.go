components {
  id: "player"
  component: "/main/assets/scripts/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Jogador.png\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tiles/player.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1.0
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
  "mask: \"vilao\"\n"
  "mask: \"item\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"esfera\"\n"
  "  }\n"
  "  data: 7.594761\n"
  "}\n"
  ""
}
