<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Stategate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Objects is an array of Object
 *
 * Generated from protobuf message <code>stategate.Objects</code>
 */
class Objects extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .stategate.Object objects = 1;</code>
     */
    private $objects;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Stategate\Object[]|\Google\Protobuf\Internal\RepeatedField $objects
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .stategate.Object objects = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getObjects()
    {
        return $this->objects;
    }

    /**
     * Generated from protobuf field <code>repeated .stategate.Object objects = 1;</code>
     * @param \Stategate\Object[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setObjects($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Stategate\Object::class);
        $this->objects = $arr;

        return $this;
    }

}

