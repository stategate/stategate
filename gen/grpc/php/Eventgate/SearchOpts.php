<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Eventgate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * SearchOpts are options when querying historical events for a given object
 *
 * Generated from protobuf message <code>eventgate.SearchOpts</code>
 */
class SearchOpts extends \Google\Protobuf\Internal\Message
{
    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
     */
    private $type = '';
    /**
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     */
    private $key = '';
    /**
     * only return events that occurred after specified min timestamp
     *
     * Generated from protobuf field <code>int64 min = 3;</code>
     */
    private $min = 0;
    /**
     * only return events that occurred before specified max timestamp
     *
     * Generated from protobuf field <code>int64 max = 4;</code>
     */
    private $max = 0;
    /**
     * limit returned events
     *
     * Generated from protobuf field <code>int64 limit = 5 [(.validator.field) = {</code>
     */
    private $limit = 0;
    /**
     * offset returned events(pagination)
     *
     * Generated from protobuf field <code>int64 offset = 6;</code>
     */
    private $offset = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $type
     *           Object type (ex: user)
     *     @type string $key
     *           Object key (unique within type)
     *     @type int|string $min
     *           only return events that occurred after specified min timestamp
     *     @type int|string $max
     *           only return events that occurred before specified max timestamp
     *     @type int|string $limit
     *           limit returned events
     *     @type int|string $offset
     *           offset returned events(pagination)
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setType($var)
    {
        GPBUtil::checkString($var, True);
        $this->type = $var;

        return $this;
    }

    /**
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getKey()
    {
        return $this->key;
    }

    /**
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->key = $var;

        return $this;
    }

    /**
     * only return events that occurred after specified min timestamp
     *
     * Generated from protobuf field <code>int64 min = 3;</code>
     * @return int|string
     */
    public function getMin()
    {
        return $this->min;
    }

    /**
     * only return events that occurred after specified min timestamp
     *
     * Generated from protobuf field <code>int64 min = 3;</code>
     * @param int|string $var
     * @return $this
     */
    public function setMin($var)
    {
        GPBUtil::checkInt64($var);
        $this->min = $var;

        return $this;
    }

    /**
     * only return events that occurred before specified max timestamp
     *
     * Generated from protobuf field <code>int64 max = 4;</code>
     * @return int|string
     */
    public function getMax()
    {
        return $this->max;
    }

    /**
     * only return events that occurred before specified max timestamp
     *
     * Generated from protobuf field <code>int64 max = 4;</code>
     * @param int|string $var
     * @return $this
     */
    public function setMax($var)
    {
        GPBUtil::checkInt64($var);
        $this->max = $var;

        return $this;
    }

    /**
     * limit returned events
     *
     * Generated from protobuf field <code>int64 limit = 5 [(.validator.field) = {</code>
     * @return int|string
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * limit returned events
     *
     * Generated from protobuf field <code>int64 limit = 5 [(.validator.field) = {</code>
     * @param int|string $var
     * @return $this
     */
    public function setLimit($var)
    {
        GPBUtil::checkInt64($var);
        $this->limit = $var;

        return $this;
    }

    /**
     * offset returned events(pagination)
     *
     * Generated from protobuf field <code>int64 offset = 6;</code>
     * @return int|string
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * offset returned events(pagination)
     *
     * Generated from protobuf field <code>int64 offset = 6;</code>
     * @param int|string $var
     * @return $this
     */
    public function setOffset($var)
    {
        GPBUtil::checkInt64($var);
        $this->offset = $var;

        return $this;
    }

}

